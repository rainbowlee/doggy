package main

import (
	"bufio"
	"strconv"
	"container/list"
	"fmt"
	"io"
	"os"
	//"errors"
	"strings"
)

func processtypeline(typeline string) (coltypes []string) {
	coltypes = strings.Split(typeline, "\t")

	//转化为go专属类型
	/*
		for index := range coltypes {
			fmt.Println(coltypes[index])
		}

		for index, value := range coltypes {
			fmt.Println(value, index)
		}
	*/

	datattypes := []string{}
	for i := 0; i < len(coltypes); i++ {
		coltype := coltypes[i]
		if coltype == "INT" {
			datattypes = append(datattypes, "int")
		} else if coltype == "INT64" {
			datattypes = append(datattypes, "int64")
		} else if coltype == "DOUBLE" {
			datattypes = append(datattypes, "float64")
		} else if coltype == "FLOAT" {
			datattypes = append(datattypes, "float32")
		} else if coltype == "string" {
			datattypes = append(datattypes, "string")
        } else if coltype == "LANGUAGE" {
			datattypes = append(datattypes, "string")
		}
        
	}

	return datattypes
}

func processidline(idline string) (colids []string) {
	colids = strings.Split(idline, "\t")
	return
}


func generate(name string, typeline string, idline string) {
	a := strings.Split(name, ".")
	configFileNameSrc := a[0]
	configFileName := strings.ToLower(a[0])
	fmt.Printf("generate %s \n %s \n %s \n %s \n", name, typeline, idline, configFileName)

	configCalss := "config" + configFileName

	fi, error := os.Create("../config/" + configCalss + ".go")
	if error != nil {
		panic(error)
	}
	defer fi.Close()

	w := bufio.NewWriter(fi)
	w.WriteString("package config \n")

	w.WriteString("\n")

	w.WriteString("import (\n")
	w.WriteString("\t\"io\"\n")
	w.WriteString("\t\"bufio\"\n")
	w.WriteString("\t\"fmt\"\n")
	w.WriteString("\t\"os\"\n")
	w.WriteString("\t\"strings\"\n")

	w.WriteString(")\n")
	w.WriteString("\n")
	w.Flush()
	types := processtypeline(typeline)
	ids := processidline(idline)
	fmt.Println(types)
	fmt.Println(ids)

	if len(types) > len(ids) {
		fmt.Printf("errorrrrrrrrrrrrrrrrrrrr!!!\n generate %s \n %s \n %s \n %s \n", name, typeline, idline, configFileName)
		return
	}

	w.WriteString("var typedefs" + configFileNameSrc + "=\"" + strings.Join(types,"\t") + "\"\n")
	w.WriteString("var iddefs" + configFileNameSrc + "=\"" + strings.Join(ids,"\t") + "\"\n")	
	w.WriteString("\n")

	//生成配置类
	w.WriteString("type " + configFileNameSrc + " struct {\n")
	for i := 0; i < len(types); i++ {
		w.WriteString("\t   " + ids[i] + "\t\t" + types[i] + "\n")
	}
	w.WriteString("}\n")
	w.WriteString("\n")
	w.Flush()

	//生成管理类
	w.WriteString("type " + configFileNameSrc + "Mgr" + " struct {\n")
	w.WriteString("\t   mapdata " + "\t\t" + "map[" + types[0] + "] " + configFileNameSrc + "\n")
	w.WriteString("}\n")
	w.WriteString("\n")

	//管理方法 load unload
	readConfigFile := "config/" + configFileNameSrc + ".txt"
	w.WriteString("func (config *" + configFileNameSrc + "Mgr) Load() {\n")
	w.WriteString("\tfi, err := os.Open(\"" + readConfigFile + "\")\n")
	w.WriteString("\tif err != nil {\n")
	w.WriteString("\t\tpanic(err)\n")
	w.WriteString("\t}\n")
	w.WriteString("\tdefer fi.Close()\n")

	w.WriteString("\tvar step = 0\n")

	w.WriteString("\tr := bufio.NewReader(fi)\n")

	loadcontent := `
	datatypes := strings.Split(typedefs%s, "\t")	
	dataids := strings.Split(iddefs%s, "\t")				
	for {
		buf, isPrefix, err := r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}

		if isPrefix == true {

		}
	
		if buf == nil {
			break
		}

		lineContent := string(buf)
		fmt.Println(lineContent)
	
		if buf[0] == '#' {
			continue
		}

		if step == 0 {
			step = 1
			continue
		}

		if step == 1 {
			step = 2
			continue
		}

		if step == 2 {
		    config.loadOneLine(lineContent, datatypes, dataids)
		}
		//chunks = append(chunks, buf...)
	}`

	loadcontent = fmt.Sprintf(loadcontent,configFileNameSrc,configFileNameSrc)
	w.WriteString("\t" + loadcontent + "\n")
	w.WriteString("}\n\n")

	online1 := `func (*%s) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(%s)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %s \n", dataline)
		return
    }`

	online1 = fmt.Sprintf(online1, configFileNameSrc + "Mgr", configFileNameSrc)
	w.WriteString(online1 + "\n")
	for index := 0; index < len(types); index++ {
		if types[index] == "int" {
			w.WriteString("\t" + "datastruct." + ids[index] + "=getIntValue(datacols[" + strconv.Itoa(index) + "])" + "\n")
		}
		if types[index] == "int64" {
			w.WriteString("\t" + "datastruct." + ids[index] + "=getInt64Value(datacols[" + strconv.Itoa(index) + "])" + "\n")
		}
		if types[index] == "string" {
			w.WriteString("\t" + "datastruct." + ids[index] + "=datacols[" + strconv.Itoa(index) + "]" + "\n")
        }
        
        if types[index] == "float64" {
			w.WriteString("\t" + "datastruct." + ids[index] + "=getFloat64Value(datacols[" + strconv.Itoa(index) + "])" + "\n")
        } 
        
        if types[index] == "float32" {
			w.WriteString("\t" + "datastruct." + ids[index] + "=getFloat32Value(datacols[" + strconv.Itoa(index) + "])" + "\n")
		}  
	}

	w.WriteString("}\n")
	w.Flush()
}

func generateCode(name string) {
	fmt.Print("generateCode  ", name)
	fi, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	var step = 0

	r := bufio.NewReader(fi)

	lineType := ""
	lineId := ""

	for {
		buf, isPrefix, err := r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}

		if isPrefix == true {

		}

		lineContent := string(buf)
		fmt.Println(lineContent)
		if buf[0] == '#' {
			continue
		}

		if step == 0 {
			lineType = lineContent
			step = 1
			continue
		}

		if step == 1 {
			lineId = lineContent
			step = 2
			break
		}

		if buf == nil {
			break
		}

		//chunks = append(chunks, buf...)
	}

	generate(name, lineType, lineId)
}

func main() {
	l := list.New()
	f, err := os.Open(".")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		fmt.Print(err)
	}
	var i = 1
	for _, name := range names {
		a := strings.Split(name, ".")
		fileExt := strings.ToLower(a[len(a)-1])
		fmt.Printf(" index %d name %s \n", i, fileExt)
		if fileExt == "txt" {
			l.PushBack(name)
			generateCode(name)
		}
		i++
	}
}

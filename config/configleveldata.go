package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsLevelData="int	int64"
var iddefsLevelData="Id	Exp"

type LevelData struct {
	   Id		int
	   Exp		int64
}

type LevelDataMgr struct {
	   mapdata 		map[int] *LevelData
}

func (config *LevelDataMgr) Load() {
	config.mapdata = make(map[int] *LevelData)
	fi, err := os.Open("config/LevelData.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	var step = 0
	r := bufio.NewReader(fi)
	
	datatypes := strings.Split(typedefsLevelData, "\t")	
	dataids := strings.Split(iddefsLevelData, "\t")				
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
	}
}

func (config *LevelDataMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(LevelData)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.Id=getIntValue(datacols[0])
	datastruct.Exp=getInt64Value(datacols[1])
	config.mapdata[datastruct.Id]=datastruct
}

func (config *LevelDataMgr) UnLoad() {
}

func (config *LevelDataMgr) GetConfig( id int) *LevelData{
	 data, ok := config.mapdata[id]
	 if ok != true {
		 return nil
	 }
	 return data
}

var LevelDataMgrInst = &LevelDataMgr{}


package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRoleData="int	string	int	int	int"
var iddefsRoleData="Id	Name	Sex	Defaulttext	Love_itemid	Max_level	Cloth_id	Qiyue_costid	Qiyuecost_num"

type RoleData struct {
	   Id		int
	   Name		string
	   Sex		int
	   Defaulttext		int
	   Love_itemid		int
}

type RoleDataMgr struct {
	   mapdata 		map[int] *RoleData
}

func (config *RoleDataMgr) Load() {
	config.mapdata = make(map[int] *RoleData)
	fi, err := os.Open("config/RoleData.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	var step = 0
	r := bufio.NewReader(fi)
	
	datatypes := strings.Split(typedefsRoleData, "\t")	
	dataids := strings.Split(iddefsRoleData, "\t")				
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

func (config *RoleDataMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RoleData)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.Id=getIntValue(datacols[0])
	datastruct.Name=datacols[1]
	datastruct.Sex=getIntValue(datacols[2])
	datastruct.Defaulttext=getIntValue(datacols[3])
	datastruct.Love_itemid=getIntValue(datacols[4])
	config.mapdata[datastruct.Id]=datastruct
}

func (config *RoleDataMgr) UnLoad() {
}

func (config *RoleDataMgr) GetConfig( id int) *RoleData{
	 data, ok := config.mapdata[id]
	 if ok != true {
		 return nil
	 }
	 return data
}

var RoleDataMgrInst = &RoleDataMgr{}


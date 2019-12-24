package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRoleData="int	string	int	int	int"
var iddefsRoleData="id	name	sex	defaulttext	love_itemid	max_level	cloth_id	qiyue_costid	qiyuecost_num"

type RoleData struct {
	   id		int
	   name		string
	   sex		int
	   defaulttext		int
	   love_itemid		int
}

type RoleDataMgr struct {
	   mapdata 		map[int] RoleData
}

func (config *RoleDataMgr) Load() {
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

func (*RoleDataMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RoleData)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.id=getIntValue(datacols[0])
	datastruct.name=datacols[1]
	datastruct.sex=getIntValue(datacols[2])
	datastruct.defaulttext=getIntValue(datacols[3])
	datastruct.love_itemid=getIntValue(datacols[4])
}

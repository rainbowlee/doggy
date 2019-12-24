package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRoleLevelData="int	int"
var iddefsRoleLevelData="id	next_exp"

type RoleLevelData struct {
	   id		int
	   next_exp		int
}

type RoleLevelDataMgr struct {
	   mapdata 		map[int] RoleLevelData
}

func (config *RoleLevelDataMgr) Load() {
	fi, err := os.Open("config/RoleLevelData.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	var step = 0
	r := bufio.NewReader(fi)
	
	datatypes := strings.Split(typedefsRoleLevelData, "\t")	
	dataids := strings.Split(iddefsRoleLevelData, "\t")				
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

func (*RoleLevelDataMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RoleLevelData)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.id=getIntValue(datacols[0])
	datastruct.next_exp=getIntValue(datacols[1])
}

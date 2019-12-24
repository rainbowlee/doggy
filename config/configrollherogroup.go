package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRollHeroGroup="int	int	int	int	int	int	int	int	int	int	int"
var iddefsRollHeroGroup="Id	Group_1	Group_2	Group_3	Group_4	Group_5	Group_6	Group_7	Group_8	Group_9	Group_10	Display	Des_l"

type RollHeroGroup struct {
	   Id		int
	   Group_1		int
	   Group_2		int
	   Group_3		int
	   Group_4		int
	   Group_5		int
	   Group_6		int
	   Group_7		int
	   Group_8		int
	   Group_9		int
	   Group_10		int
}

type RollHeroGroupMgr struct {
	   mapdata 		map[int] *RollHeroGroup
}

func (config *RollHeroGroupMgr) Load() {
	config.mapdata = make(map[int] *RollHeroGroup)
	fi, err := os.Open("config/RollHeroGroup.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	var step = 0
	r := bufio.NewReader(fi)
	
	datatypes := strings.Split(typedefsRollHeroGroup, "\t")	
	dataids := strings.Split(iddefsRollHeroGroup, "\t")				
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

func (config *RollHeroGroupMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RollHeroGroup)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.Id=getIntValue(datacols[0])
	datastruct.Group_1=getIntValue(datacols[1])
	datastruct.Group_2=getIntValue(datacols[2])
	datastruct.Group_3=getIntValue(datacols[3])
	datastruct.Group_4=getIntValue(datacols[4])
	datastruct.Group_5=getIntValue(datacols[5])
	datastruct.Group_6=getIntValue(datacols[6])
	datastruct.Group_7=getIntValue(datacols[7])
	datastruct.Group_8=getIntValue(datacols[8])
	datastruct.Group_9=getIntValue(datacols[9])
	datastruct.Group_10=getIntValue(datacols[10])
	config.mapdata[datastruct.Id]=datastruct
}

func (config *RollHeroGroupMgr) UnLoad() {
}

func (config *RollHeroGroupMgr) GetConfig( id int) *RollHeroGroup{
	 data, ok := config.mapdata[id]
	 if ok != true {
		 return nil
	 }
	 return data
}

var RollHeroGroupMgrInst = &RollHeroGroupMgr{}


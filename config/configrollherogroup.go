package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRollHeroGroup="int	int	int	int	int	int	int	int	int	int	int"
var iddefsRollHeroGroup="id	group_1	group_2	group_3	group_4	group_5	group_6	group_7	group_8	group_9	group_10	display	des_l"

type RollHeroGroup struct {
	   id		int
	   group_1		int
	   group_2		int
	   group_3		int
	   group_4		int
	   group_5		int
	   group_6		int
	   group_7		int
	   group_8		int
	   group_9		int
	   group_10		int
}

type RollHeroGroupMgr struct {
	   mapdata 		map[int] RollHeroGroup
}

func (config *RollHeroGroupMgr) Load() {
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

func (*RollHeroGroupMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RollHeroGroup)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.id=getIntValue(datacols[0])
	datastruct.group_1=getIntValue(datacols[1])
	datastruct.group_2=getIntValue(datacols[2])
	datastruct.group_3=getIntValue(datacols[3])
	datastruct.group_4=getIntValue(datacols[4])
	datastruct.group_5=getIntValue(datacols[5])
	datastruct.group_6=getIntValue(datacols[6])
	datastruct.group_7=getIntValue(datacols[7])
	datastruct.group_8=getIntValue(datacols[8])
	datastruct.group_9=getIntValue(datacols[9])
	datastruct.group_10=getIntValue(datacols[10])
}

package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRollHeroDrop="int	int	int	int	int	int	int	int	int	int"
var iddefsRollHeroDrop="id	group_id	typeid	param	count	seed	lucky_add	lucky_full	seed_change	combine_group"

type RollHeroDrop struct {
	   id		int
	   group_id		int
	   typeid		int
	   param		int
	   count		int
	   seed		int
	   lucky_add		int
	   lucky_full		int
	   seed_change		int
	   combine_group		int
}

type RollHeroDropMgr struct {
	   mapdata 		map[int] RollHeroDrop
}

func (config *RollHeroDropMgr) Load() {
	fi, err := os.Open("config/RollHeroDrop.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	var step = 0
	r := bufio.NewReader(fi)
	
	datatypes := strings.Split(typedefsRollHeroDrop, "\t")	
	dataids := strings.Split(iddefsRollHeroDrop, "\t")				
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

func (*RollHeroDropMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RollHeroDrop)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.id=getIntValue(datacols[0])
	datastruct.group_id=getIntValue(datacols[1])
	datastruct.typeid=getIntValue(datacols[2])
	datastruct.param=getIntValue(datacols[3])
	datastruct.count=getIntValue(datacols[4])
	datastruct.seed=getIntValue(datacols[5])
	datastruct.lucky_add=getIntValue(datacols[6])
	datastruct.lucky_full=getIntValue(datacols[7])
	datastruct.seed_change=getIntValue(datacols[8])
	datastruct.combine_group=getIntValue(datacols[9])
}

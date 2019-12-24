package config 

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typedefsRollHeroDrop="int	int	int	int	int	int	int	int	int	int"
var iddefsRollHeroDrop="Id	Group_id	Typeid	Param	Count	Seed	Lucky_add	Lucky_full	Seed_change	Combine_group"

type RollHeroDrop struct {
	   Id		int
	   Group_id		int
	   Typeid		int
	   Param		int
	   Count		int
	   Seed		int
	   Lucky_add		int
	   Lucky_full		int
	   Seed_change		int
	   Combine_group		int
}

type RollHeroDropMgr struct {
	   mapdata 		map[int] *RollHeroDrop
}

func (config *RollHeroDropMgr) Load() {
	config.mapdata = make(map[int] *RollHeroDrop)
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

func (config *RollHeroDropMgr) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(RollHeroDrop)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %!s(MISSING) \n", dataline)
		return
    }
	datastruct.Id=getIntValue(datacols[0])
	datastruct.Group_id=getIntValue(datacols[1])
	datastruct.Typeid=getIntValue(datacols[2])
	datastruct.Param=getIntValue(datacols[3])
	datastruct.Count=getIntValue(datacols[4])
	datastruct.Seed=getIntValue(datacols[5])
	datastruct.Lucky_add=getIntValue(datacols[6])
	datastruct.Lucky_full=getIntValue(datacols[7])
	datastruct.Seed_change=getIntValue(datacols[8])
	datastruct.Combine_group=getIntValue(datacols[9])
	config.mapdata[datastruct.Id]=datastruct
}

func (config *RollHeroDropMgr) UnLoad() {
}

func (config *RollHeroDropMgr) GetConfig( id int) *RollHeroDrop{
	 data, ok := config.mapdata[id]
	 if ok != true {
		 return nil
	 }
	 return data
}

var RollHeroDropMgrInst = &RollHeroDropMgr{}


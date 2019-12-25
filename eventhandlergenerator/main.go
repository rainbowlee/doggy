package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	//"errors"
	"doggy/common"
	"strings"
)

// messagetype & messagemap MessageIdBase MessageIdHeader
var (
	MessageType     = []string{"Dse", "Dce", "Drd", "Ddr", "Dsr", "Drs", "Dsw", "Dws"}
	MessageIdBase   = []int{1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000}
	MessageIdHeader = []string{"EVENT_SE", "EVENT_CE", "EVENT_RD", "EVENT_DR", "EVENT_SR", "EVENT_RS", "EVENT_SW", "EVENT_WS"}

	MessageDefs    = make(map[string]*list.List) //最新配置文件
	MessageDefOlds = make(map[string]*list.List) //原有定义 EventId文件读取

	MessageMap    = make(map[string]string)
	MessageMapOld = make(map[string]string)

	MessageHandler = []string{}
)

func getMessageTypeAndIndex(message string) (index int) {
	for index, value := range MessageType {
		if message[:3] == value {
			return index
		}
	}

	return -1
}

func getMessageIdFromEventIdLine(line string) (messageid string) {
	for _, value := range MessageIdHeader {
		findindex := strings.Index(line, value)
		if findindex != -1 {
			findindex2 := strings.Index(line, " = ")

			messageid = line[findindex+len(MessageIdHeader[0])+1 : findindex2] //多一个—
			return messageid
		}
	}

	return ""
}

func dellist(l *list.List, value string) {
	for iter := l.Front(); iter != nil; iter.Next() {
		if iter.Value == value {
			l.Remove(iter)
			return
		}
	}
}

func genEventFile() {
	fi, error := os.Create("../eventhandler/EventId.go")
	if error != nil {
		panic(error)
	}
	defer fi.Close()

	w := bufio.NewWriter(fi)
	w.WriteString("package eventhandler \n")
	w.WriteString("\n")

	for index, typevalue := range MessageType {
		w.WriteString("//event id for " + typevalue + "\n")
		w.WriteString("const (\n")

		MessageHandler = []string{}
		MessageIdHeader := MessageIdHeader[index]
		BaseId := MessageIdBase[index]

		l, ok := MessageDefOlds[typevalue]
		if ok == true {
			for i := l.Front(); i != nil; i = i.Next() {
				valuestr, _ := i.Value.(string)
				_, ok := MessageMap[valuestr]
				if ok == false {
					continue
				}
				genOneEventHandler(typevalue, valuestr, MessageIdHeader+"_"+valuestr)
				outvalue := fmt.Sprintf("\t"+MessageIdHeader+"_%s = %d\n", i.Value, BaseId)
				BaseId = BaseId + 1
				w.WriteString(outvalue)
			}
		}

		l, ok = MessageDefs[typevalue]
		if ok == true {
			for i := l.Front(); i != nil; i = i.Next() {
				valuestr, _ := i.Value.(string)
				_, ok := MessageMapOld[valuestr]
				if ok == true {
					continue
				}
				genOneEventHandler(typevalue, valuestr, MessageIdHeader+"_"+valuestr)
				outvalue := fmt.Sprintf("\t"+MessageIdHeader+"_%s = %d\n", i.Value, BaseId)
				BaseId = BaseId + 1
				w.WriteString(outvalue)
			}
		}

		w.WriteString(")\n\n\n")

		genEventDispatcher(typevalue)
	}

	w.Flush()
}

func genEventDispatcher(messageType string) {
	fmt.Println("genEventDispatcher", messageType)
	name := "../eventhandler/" + messageType + "EventDispatcher.go"
	fi, error := os.Create(name)
	if error != nil {
		panic(error)
	}
	defer fi.Close()

	w := bufio.NewWriter(fi)
	w.WriteString("package eventhandler \n")
	w.WriteString("\n")
	w.WriteString("import (\n")
	w.WriteString("\t\"doggy/common\"\n")
//	w.WriteString("\t\"fmt\"\n")	
//	w.WriteString("\t\"strconv\"\n")	
	w.WriteString(")\n")
	w.WriteString("\n")

	dipatchertext := `
type %sEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *%sEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *%sEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *%sEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func %sEventDispatcherInst() *%sEventDispatcher{
	return &dispatcher%s
}

var dispatcher%s = %sEventDispatcher{}
`
	dipatchertext = fmt.Sprintf(dipatchertext, messageType, messageType, messageType, messageType, messageType, messageType, messageType, messageType, messageType)
	w.WriteString(dipatchertext)
	w.WriteString("\n")

	w.Flush()
}

//messageId
func genOneEventHandler(messageType string, messageDef string, messageId string) {
	fmt.Println(messageType, messageDef, messageId)
	name := "../eventhandler/" + messageDef + "Handler.go"
	common.FileExist(name)

	//eventhandler :=new(common.EventHandler)
	//eventhandler.Handle()

	fi, error := os.Create(name)
	if error != nil {
		panic(error)
	}
	defer fi.Close()

	w := bufio.NewWriter(fi)
	w.WriteString("package eventhandler \n")
	w.WriteString("\n")
	w.WriteString("import (\n")
	w.WriteString("\t\"doggy/common\"\n")
	w.WriteString(")\n")
	w.WriteString("\n")

	handlerText := `
type %s struct {
}

func (eventhandler *%s) Register() {
	%s
}

func (eventhandler *%s) Handle(event*common.Event) {

}`

	param1 := messageDef + "Handler"
	param2 := messageType + "EventDispatcherInst().RegisterEventHandler(" + messageId + ", (*common.EventHandler)(eventhandler))"

	handlerText = fmt.Sprintf(handlerText, param1, param1,param2,param1)
	w.WriteString(handlerText)
	w.Flush()
	MessageHandler = append(MessageHandler, param1)
}

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
	var retcolids = []string{}
	for _, value := range colids {
		char0 := value[0] - 32
		newvalue := string(char0) + value[1:]
		retcolids = append(retcolids, newvalue)
	}
	return retcolids
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

	w.WriteString("var typedefs" + configFileNameSrc + "=\"" + strings.Join(types, "\t") + "\"\n")
	w.WriteString("var iddefs" + configFileNameSrc + "=\"" + strings.Join(ids, "\t") + "\"\n")
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
	w.WriteString("\t   mapdata " + "\t\t" + "map[" + types[0] + "] *" + configFileNameSrc + "\n")
	w.WriteString("}\n")
	w.WriteString("\n")

	//管理方法 load unload
	readConfigFile := "config/" + configFileNameSrc + ".txt"
	w.WriteString("func (config *" + configFileNameSrc + "Mgr) Load() {\n")
	w.WriteString("\tconfig.mapdata = make(map[" + types[0] + "] *" + configFileNameSrc + ")\n")

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

	loadcontent = fmt.Sprintf(loadcontent, configFileNameSrc, configFileNameSrc)
	w.WriteString("\t" + loadcontent + "\n")
	w.WriteString("}\n\n")

	online1 := `func (config *%s) loadOneLine(dataline string, datatypes []string, dataids []string) {
	datacols := strings.Split(dataline, "\t")
	datastruct := new(%s)
	if len(datatypes) != len(datacols) {
		fmt.Printf("read data error %s \n", dataline)
		return
    }`

	online1 = fmt.Sprintf(online1, configFileNameSrc+"Mgr", configFileNameSrc)
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
	w.WriteString("\tconfig.mapdata[datastruct." + ids[0] + "]=" + "datastruct\n")
	w.WriteString("}\n")

	w.WriteString("\n")
	//unload
	w.WriteString("func (config *" + configFileNameSrc + "Mgr) UnLoad() {\n")
	w.WriteString("}\n")
	w.WriteString("\n")

	w.WriteString("func (config *" + configFileNameSrc + "Mgr) GetConfig( id " + types[0] + ") *" + configFileNameSrc + "{\n")
	w.WriteString("\t data, ok := config.mapdata[id]\n")
	w.WriteString("\t if ok != true {\n")
	w.WriteString("\t\t return nil\n")
	w.WriteString("\t }\n")
	w.WriteString("\t return data\n")
	w.WriteString("}\n")
	w.WriteString("\n")

	w.WriteString("var " + configFileNameSrc + "MgrInst = &" + configFileNameSrc + "Mgr{}\r\n")
	w.WriteString("\n")
	w.Flush()
}

func readeventidfile() {
	name := "../eventhandler/EventId.go"
	fmt.Println("readeventidfile  ", name)
	fi, err := os.Open(name)
	if err != nil {
		return
	}

	defer fi.Close()
	r := bufio.NewReader(fi)

	for {
		buf, isPrefix, err := r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}

		if isPrefix == true {

		}

		lineContent := string(buf)

		if buf == nil {
			break
		}

		if len(buf) == 0 {
			continue
		}

		if buf[0] == '#' {
			continue
		}

		//message 后一个空格
		messsageid := getMessageIdFromEventIdLine(lineContent)
		if messsageid != "" {
			typeindex := getMessageTypeAndIndex(messsageid)
			messagetype := MessageType[typeindex]
			l, ok := MessageDefOlds[messagetype]
			if ok == false {
				l = list.New()
				MessageDefOlds[messagetype] = l
			}
			l.PushBack(messsageid)
			MessageMapOld[messsageid] = messsageid
		}
		fmt.Println(messsageid)
	}
}

func readprotofile(name string) {
	fmt.Println("readprotofile  ", name)
	fi, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer fi.Close()
	r := bufio.NewReader(fi)

	for {
		buf, isPrefix, err := r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}

		if isPrefix == true {

		}

		lineContent := string(buf)

		if buf == nil {
			break
		}

		if len(buf) == 0 {
			continue
		}

		if buf[0] == '#' {
			continue
		}

		//message 后一个空格
		if len(lineContent) > 7 && lineContent[0:7] == "message" {
			message := lineContent[8:]

			index2 := strings.Index(message, "\t")
			index3 := strings.Index(message, " ")

			index := len(message)
			if index2 != -1 && index > index2 {
				index = index2
			}

			if index3 != -1 && index > index3 {
				index = index3
			}

			messagedef := message[:index]

			messagetypeindex := getMessageTypeAndIndex(messagedef)
			if messagetypeindex == -1 {
				fmt.Println("error messagedef:", messagedef)
			} else {
				messgetype := MessageType[messagetypeindex]
				//v2, ok2 := m["x"]
				typelist, exists := MessageDefs[messgetype]
				if exists == false {
					typelist = list.New()
					MessageDefs[messgetype] = typelist
				}

				typelist.PushBack(messagedef)
				MessageMap[messagedef] = messagedef
			}
		}
	}
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
		if fileExt == "proto" {
			l.PushBack(name)
			readprotofile(name)
		}
		i++
	}
	readeventidfile()
	genEventFile()
}

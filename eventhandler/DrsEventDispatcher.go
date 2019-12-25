package eventhandler 

import (
	"doggy/common"
)


type DrsEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DrsEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DrsEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DrsEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DrsEventDispatcherInst() *DrsEventDispatcher{
	return &dispatcherDrs
}

var dispatcherDrs = DrsEventDispatcher{}


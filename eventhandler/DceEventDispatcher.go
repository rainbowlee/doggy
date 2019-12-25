package eventhandler 

import (
	"doggy/common"
)


type DceEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DceEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DceEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DceEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DceEventDispatcherInst() *DceEventDispatcher{
	return &dispatcherDce
}

var dispatcherDce = DceEventDispatcher{}


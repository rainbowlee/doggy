package eventhandler 

import (
	"doggy/common"
)


type DswEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DswEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DswEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DswEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DswEventDispatcherInst() *DswEventDispatcher{
	return &dispatcherDsw
}

var dispatcherDsw = DswEventDispatcher{}


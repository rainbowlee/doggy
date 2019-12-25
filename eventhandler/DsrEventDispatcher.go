package eventhandler 

import (
	"doggy/common"
)


type DsrEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DsrEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DsrEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DsrEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DsrEventDispatcherInst() *DsrEventDispatcher{
	return &dispatcherDsr
}

var dispatcherDsr = DsrEventDispatcher{}


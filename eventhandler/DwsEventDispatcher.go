package eventhandler 

import (
	"doggy/common"
)


type DwsEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DwsEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DwsEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DwsEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DwsEventDispatcherInst() *DwsEventDispatcher{
	return &dispatcherDws
}

var dispatcherDws = DwsEventDispatcher{}


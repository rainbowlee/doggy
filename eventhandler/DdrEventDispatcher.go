package eventhandler 

import (
	"doggy/common"
)


type DdrEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DdrEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DdrEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DdrEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DdrEventDispatcherInst() *DdrEventDispatcher{
	return &dispatcherDdr
}

var dispatcherDdr = DdrEventDispatcher{}


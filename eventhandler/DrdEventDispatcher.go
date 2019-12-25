package eventhandler 

import (
	"doggy/common"
)


type DrdEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DrdEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DrdEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DrdEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DrdEventDispatcherInst() *DrdEventDispatcher{
	return &dispatcherDrd
}

var dispatcherDrd = DrdEventDispatcher{}


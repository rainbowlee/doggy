package eventhandler 

import (
	"doggy/common"
)


type DseEventDispatcher struct {
	common.EventDispatcher
}
/*
func (dispatcher *DseEventDispatcher) RegisterEventHandler(eventid int, eventhandler* common.EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *DseEventDispatcher) Dispach(event *common.Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *DseEventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
*/

func DseEventDispatcherInst() *DseEventDispatcher{
	return &dispatcherDse
}

var dispatcherDse = DseEventDispatcher{}


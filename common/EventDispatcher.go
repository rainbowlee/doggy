package common

import (
	"fmt"
	"strconv"
)

type EventDispatcher struct {
	mapHandlers map[int]*EventHandler
}

func (dispatcher *EventDispatcher) RegisterEventHandler(eventid int, eventhandler*EventHandler) {
	dispatcher.mapHandlers[eventid] = eventhandler
}

func (dispatcher *EventDispatcher) Dispach(event*Event) {
	handlerdef, ok :=dispatcher.mapHandlers[event.EventId]
	if ok == true{
		handlerdef.Handle(event)
	} else{
		fmt.Println("undefied msg " + strconv.Itoa(event.EventId) + "\n")
	}
}

func (dispatcher *EventDispatcher) Init(){
	dispatcher.mapHandlers = make(map[int]*EventHandler)
}
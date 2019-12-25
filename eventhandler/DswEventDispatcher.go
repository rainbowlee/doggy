package eventhandler 

import (
	"doggy/common"
)


type DswEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DswEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	

}

func DswEventDispatcherInst() *DswEventDispatcher{
	return &dispatcherDsw
}

var dispatcherDsw = DswEventDispatcher{}


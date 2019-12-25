package eventhandler 

import (
	"doggy/common"
)


type DrsEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DrsEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	

}

func DrsEventDispatcherInst() *DrsEventDispatcher{
	return &dispatcherDrs
}

var dispatcherDrs = DrsEventDispatcher{}


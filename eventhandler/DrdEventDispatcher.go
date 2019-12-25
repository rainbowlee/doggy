package eventhandler 

import (
	"doggy/common"
)


type DrdEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DrdEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	

}

func DrdEventDispatcherInst() *DrdEventDispatcher{
	return &dispatcherDrd
}

var dispatcherDrd = DrdEventDispatcher{}


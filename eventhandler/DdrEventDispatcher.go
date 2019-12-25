package eventhandler 

import (
	"doggy/common"
)


type DdrEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DdrEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	

}

func DdrEventDispatcherInst() *DdrEventDispatcher{
	return &dispatcherDdr
}

var dispatcherDdr = DdrEventDispatcher{}


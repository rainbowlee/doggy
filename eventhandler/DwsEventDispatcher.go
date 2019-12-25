package eventhandler 

import (
	"doggy/common"
)


type DwsEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DwsEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	

}

func DwsEventDispatcherInst() *DwsEventDispatcher{
	return &dispatcherDws
}

var dispatcherDws = DwsEventDispatcher{}


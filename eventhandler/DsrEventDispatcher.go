package eventhandler 

import (
	"doggy/common"
)


type DsrEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DsrEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	

}

func DsrEventDispatcherInst() *DsrEventDispatcher{
	return &dispatcherDsr
}

var dispatcherDsr = DsrEventDispatcher{}


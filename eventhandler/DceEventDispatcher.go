package eventhandler 

import (
	"doggy/common"
)


type DceEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DceEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	
	new(DceAuthHandler).Register()
	new(DceCardGroupRenameHandler).Register()
	new(DceCardSetMainCardHandler).Register()
	new(DceCardSetCardGroupHandler).Register()

}

func DceEventDispatcherInst() *DceEventDispatcher{
	return &dispatcherDce
}

var dispatcherDce = DceEventDispatcher{}


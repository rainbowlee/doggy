package eventhandler 

import (
	"doggy/common"
)


type DseEventDispatcher struct {
	common.EventDispatcher
}

func (dispatcher *DseEventDispatcher) RegisterHandlers(){
	dispatcher.Init()
	
	new(DseAuthHandler).Register()
	new(DseCardSendDataHandler).Register()
	new(DseCardSendSingleDataHandler).Register()
	new(DseCardGroupRenameHandler).Register()
	new(DseCardSetMainCardHandler).Register()
	new(DseCardSetCardGroupHandler).Register()

}

func DseEventDispatcherInst() *DseEventDispatcher{
	return &dispatcherDse
}

var dispatcherDse = DseEventDispatcher{}


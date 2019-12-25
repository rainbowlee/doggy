package eventhandler 

import (
	"doggy/common"
)


type DceCardSetMainCardHandler struct {
}

func (eventhandler *DceCardSetMainCardHandler) Register() {
	DceEventDispatcherInst().RegisterEventHandler(EVENT_CE_DceCardSetMainCard, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DceCardSetMainCardHandler) Handle(event*common.Event) {

}
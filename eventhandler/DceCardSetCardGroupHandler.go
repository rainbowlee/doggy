package eventhandler 

import (
	"doggy/common"
)


type DceCardSetCardGroupHandler struct {
}

func (eventhandler *DceCardSetCardGroupHandler) Register() {
	DceEventDispatcherInst().RegisterEventHandler(EVENT_CE_DceCardSetCardGroup, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DceCardSetCardGroupHandler) Handle(event*common.Event) {

}
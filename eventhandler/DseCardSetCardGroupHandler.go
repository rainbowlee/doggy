package eventhandler 

import (
	"doggy/common"
)


type DseCardSetCardGroupHandler struct {
}

func (eventhandler *DseCardSetCardGroupHandler) Register() {
	DseEventDispatcherInst().RegisterEventHandler(EVENT_SE_DseCardSetCardGroup, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DseCardSetCardGroupHandler) Handle(event*common.Event) {

}
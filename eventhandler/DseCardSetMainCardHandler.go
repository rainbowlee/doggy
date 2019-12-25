package eventhandler 

import (
	"doggy/common"
)


type DseCardSetMainCardHandler struct {
}

func (eventhandler *DseCardSetMainCardHandler) Register() {
	DseEventDispatcherInst().RegisterEventHandler(EVENT_SE_DseCardSetMainCard, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DseCardSetMainCardHandler) Handle(event*common.Event) {

}
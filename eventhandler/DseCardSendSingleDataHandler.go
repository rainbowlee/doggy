package eventhandler 

import (
	"doggy/common"
)


type DseCardSendSingleDataHandler struct {
}

func (eventhandler *DseCardSendSingleDataHandler) Register() {
	DseEventDispatcherInst().RegisterEventHandler(EVENT_SE_DseCardSendSingleData, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DseCardSendSingleDataHandler) Handle(event*common.Event) {

}
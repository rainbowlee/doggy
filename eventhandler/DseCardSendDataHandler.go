package eventhandler 

import (
	"doggy/common"
)


type DseCardSendDataHandler struct {
}

func (eventhandler *DseCardSendDataHandler) Register() {
	DseEventDispatcherInst().RegisterEventHandler(EVENT_SE_DseCardSendData, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DseCardSendDataHandler) Handle(event*common.Event) {

}
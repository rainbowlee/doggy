package eventhandler 

import (
	"doggy/common"
)


type DseAuthHandler struct {
}

func (eventhandler *DseAuthHandler) Register() {
	DseEventDispatcherInst().RegisterEventHandler(EVENT_SE_DseAuth, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DseAuthHandler) Handle(event*common.Event) {

}
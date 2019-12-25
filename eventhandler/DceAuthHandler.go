package eventhandler 

import (
	"doggy/common"
)


type DceAuthHandler struct {
}

func (eventhandler *DceAuthHandler) Register() {
	DceEventDispatcherInst().RegisterEventHandler(EVENT_CE_DceAuth, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DceAuthHandler) Handle(event*common.Event) {

}
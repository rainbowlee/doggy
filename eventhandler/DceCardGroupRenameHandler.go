package eventhandler 

import (
	"doggy/common"
)


type DceCardGroupRenameHandler struct {
}

func (eventhandler *DceCardGroupRenameHandler) Register() {
	DceEventDispatcherInst().RegisterEventHandler(EVENT_CE_DceCardGroupRename, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DceCardGroupRenameHandler) Handle(event*common.Event) {

}
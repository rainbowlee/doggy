package eventhandler 

import (
	"doggy/common"
)


type DseCardGroupRenameHandler struct {
}

func (eventhandler *DseCardGroupRenameHandler) Register() {
	DseEventDispatcherInst().RegisterEventHandler(EVENT_SE_DseCardGroupRename, (*common.EventHandler)(eventhandler))
}

func (eventhandler *DseCardGroupRenameHandler) Handle(event*common.Event) {

}
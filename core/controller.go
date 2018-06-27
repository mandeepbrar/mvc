package core

import (
	"sync"

	"github.com/mandeepbrar/mvc/interfaces"
)

type controller struct {
	commandMap  map[string]interfaces.Command
	multitonKey string
	view        interfaces.View
}

var (
	ctlonce        sync.Once
	ctlInstanceMap map[string]*controller
)

func newController(key string) *controller {
	ctl := &controller{make(map[string]interfaces.Command), key, nil}
	ctlInstanceMap[key] = ctl
	ctl.InitializeController()
	return ctl
}

func (ctl *controller) InitializeController() {
	ctl.view = GetViewInstance(ctl.multitonKey)
}

func GetControllerInstance(key string) *controller {
	ctlonce.Do(func() {
		ctlInstanceMap = make(map[string]*controller)
	})
	ins, ok := ctlInstanceMap[key]
	if !ok {
		return newController(key)
	} else {
		return ins
	}
}

func RemoveController(key string) {
	delete(ctlInstanceMap, key)
}

func (ctl *controller) ExecuteCommand(note interfaces.Notification) {
	commandInstance, ok := ctl.commandMap[note.GetName()]
	if ok {
		commandInstance.InitializeNotifier(ctl.multitonKey)
		commandInstance.Execute(note)
	}
}

func (ctl *controller) RegisterCommand(notificationName string, command interfaces.Command) {
	_, ok := ctl.commandMap[notificationName]
	if ok {
		return
	} else {
		ctl.commandMap[notificationName] = command
	}
	ctl.view.RegisterObserver(notificationName, NewObserver(func(notification interfaces.Notification) {
		ctl.ExecuteCommand(notification)
	}, ctl))
}

func (ctl *controller) RemoveCommand(notificationName string) {
	// if the Command is registered...
	if ctl.HasCommand(notificationName) {
		// remove the observer
		ctl.view.RemoveObserver(notificationName, ctl)
		delete(ctl.commandMap, notificationName)
	}
}

func (ctl *controller) HasCommand(notificationName string) bool {
	_, ok := ctl.commandMap[notificationName]
	return ok
}

package patterns

import (
	"sync"

	"github.com/mandeepbrar/mvc/core"
	"github.com/mandeepbrar/mvc/interfaces"
)

type facade struct {
	multitonKey string
	controller  interfaces.Controller
	model       interfaces.Model
	view        interfaces.View
}

var (
	fonce        sync.Once
	fInstanceMap map[string]*facade
)

func newFacade(key string) *facade {
	f := &facade{multitonKey: key}
	fInstanceMap[key] = f
	f.InitializeFacade()
	return f
}

func (f *facade) InitializeFacade() {
	f.InitializeController()
	f.InitializeModel()
	f.InitializeView()
}

func GetFacadeInstance(key string) *facade {
	fonce.Do(func() {
		fInstanceMap = make(map[string]*facade)
	})
	ins, ok := fInstanceMap[key]
	if !ok {
		return newFacade(key)
	} else {
		return ins
	}
}

func HasFacade(key string) bool {
	_, ok := fInstanceMap[key]
	return ok
}

func RemoveFacade(key string) {
	core.RemoveModel(key)
	core.RemoveView(key)
	core.RemoveController(key)
	delete(fInstanceMap, key)
}

func (f *facade) InitializeController() {
	if f.controller == nil {
		f.controller = core.GetControllerInstance(f.multitonKey)
	}
}

func (f *facade) InitializeModel() {
	if f.model == nil {
		f.model = core.GetModelInstance(f.multitonKey)
	}
}

func (f *facade) InitializeView() {
	if f.view == nil {
		f.view = core.GetViewInstance(f.multitonKey)
	}
}

func (f *facade) RegisterCommand(notificationName string, command interfaces.Command) {
	f.controller.RegisterCommand(notificationName, command)
}
func (f *facade) RemoveCommand(notificationName string) {
	f.controller.RemoveCommand(notificationName)
}

func (f *facade) HasCommand(notificationName string) bool {
	return f.controller.HasCommand(notificationName)
}
func (f *facade) RegisterProxy(proxy interfaces.Proxy) {
	f.model.RegisterProxy(proxy)
}

func (f *facade) RetrieveProxy(proxyName string) interfaces.Proxy {
	return f.model.RetrieveProxy(proxyName)
}

func (f *facade) RemoveProxy(proxyName string) interfaces.Proxy {
	if f.model != nil {
		return f.model.RemoveProxy(proxyName)
	}
	return nil
}

func (f *facade) HasProxy(proxyName string) bool {
	if f.model != nil {
		return f.model.HasProxy(proxyName)
	}
	return false
}

func (f *facade) RegisterMediator(mediator interfaces.Mediator) {
	if f.view != nil {
		f.view.RegisterMediator(mediator)
	}
}

func (f *facade) RetrieveMediator(mediatorName string) interfaces.Mediator {
	return f.view.RetrieveMediator(mediatorName)
}

func (f *facade) RemoveMediator(mediatorName string) interfaces.Mediator {
	if f.view != nil {
		return f.view.RemoveMediator(mediatorName)
	}
	return nil
}

func (f *facade) HasMediator(mediatorName string) bool {
	if f.view != nil {
		return f.view.HasMediator(mediatorName)
	}
	return false
}
func (f *facade) NotifyObservers(notification interfaces.Notification) {
	if f.view != nil {
		f.view.NotifyObservers(notification)
	}
}

func (f *facade) InitializeNotifier(key string) {
	f.multitonKey = key
}

func (f *facade) SendNotification(notificationName string, body interface{}, typeName string) {
	f.NotifyObservers(&notification{notificationName, body, typeName})
}

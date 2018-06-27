package patterns

import (
	"mvc/api"
	"mvc/core"
	"sync"
)

type facade struct {
	multitonKey string
	controller api.Controller
	model api.Model
	view api.View
}

var (
	fonce sync.Once
	fInstanceMap map[string]*facade
)

func newFacade(key string) *facade {
	f := &facade{multitonKey: key}
	fInstanceMap[key] = f;
	f.InitializeFacade()
	return f
}

func (f *facade) InitializeFacade() {
	if (f.controller == nil) {
		f.controller = core.GetControllerInstance(f.multitonKey);	
	}
	if (f.model == nil) {
		f.model = core.GetModelInstance(f.multitonKey);	
	}
	if (f.view == nil) {
		f.view = core.GetViewInstance(f.multitonKey);	
	}
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
	core.RemoveModel(key);
	core.RemoveView(key);
	core.RemoveController(key);
	delete(fInstanceMap, key)
}

func (f *facade)RegisterCommand( notificationName string, command api.Command) {
	f.controller.RegisterCommand(notificationName, command);
}
func (f *facade)RemoveCommand( notificationName string) {
	f.controller.RemoveCommand(notificationName);
}

func (f *facade)HasCommand( notificationName string ) bool {
	return f.controller.HasCommand(notificationName);
}
func (f *facade)RegisterProxy( proxy api.Proxy) {
	f.model.RegisterProxy(proxy);
}

func (f *facade)RetrieveProxy(proxyName string) api.Proxy {
	return f.model.RetrieveProxy(proxyName);
}

func (f *facade)RemoveProxy(proxyName string) api.Proxy {
	if (f.model != nil) {
		return f.model.RemoveProxy(proxyName);
	}
	return nil;
}

func (f *facade)HasProxy( proxyName string) bool {
	if (f.model != nil) {
		return f.model.HasProxy(proxyName);
	}
	return false
}

func (f *facade)RegisterMediator(mediator api.Mediator){
	if (f.view != nil) {
		f.view.RegisterMediator(mediator);
	}
}

func (f *facade)RetrieveMediator(mediatorName string) api.Mediator{
	return f.view.RetrieveMediator(mediatorName);
}

func (f *facade)RemoveMediator(mediatorName string ) api.Mediator {
	if (f.view != nil) {
		return f.view.RemoveMediator(mediatorName);
	}
	return nil;
}

func (f *facade)HasMediator(mediatorName string ) bool {
	if (f.view != nil) {
		return f.view.HasMediator(mediatorName);
	}
	return false
}
func (f *facade)NotifyObservers(notification api.Notification) {
	if (f.view != nil) {
		f.view.NotifyObservers(notification);
	}
}

func (f *facade) InitializeNotifier(key string) {
	f.multitonKey = key;
}

func (f *facade) SendNotification(notificationName string, body interface{}, typeName string) {
	f.NotifyObservers(&notification{notificationName, body, typeName});
}
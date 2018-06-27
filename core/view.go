package core

import (
	"mvc/api"
	"sync"
)

type view struct {
	observerMap map[string][]api.Observer
	mediatorMap map[string]api.Mediator
	multitonKey string
}


var (
	vwonce sync.Once
	vwInstanceMap map[string]*view
)

func newView(key string) *view {
	vw := &view{make(map[string][]api.Observer), make(map[string]api.Mediator), key}
	vwInstanceMap[key] = vw;
	vw.InitializeView()
	return vw
}

func (vw *view) InitializeView() {
}


func GetViewInstance(key string) *view {
    vwonce.Do(func() {
        vwInstanceMap = make(map[string]*view)
    })
	ins, ok := vwInstanceMap[key]
	if !ok {
		return newView(key)
	} else {
		return ins
	}
}

func RemoveView(key string) {
	delete(vwInstanceMap, key)
}

func (vw *view) HasMediator(mediatorName string) bool {
	_, ok := vw.mediatorMap[mediatorName];
	return ok
}

func (vw *view) RegisterMediator(mediator api.Mediator) {
	medName := mediator.GetMediatorName()
	_, ok := vw.mediatorMap[medName]
	if (!ok) {
		mediator.InitializeNotifier(vw.multitonKey);

		vw.mediatorMap[medName] = mediator;

		noteInterests := mediator.ListNotificationInterests();

		for _, interest := range noteInterests {
			vw.RegisterObserver(interest, NewObserver(func(note api.Notification){
				mediator.HandleNotification(note)
			}, mediator))
		}

		mediator.OnRegister();
	}
}

func (vw *view) RetrieveMediator(key string) api.Mediator {
	return vw.mediatorMap[key]
}
	
func (vw *view) RemoveMediator(mediatorName string) api.Mediator {
	mediator, ok := vw.mediatorMap[mediatorName];

	if(ok) {
		interests := mediator.ListNotificationInterests();
		for _, interest := range interests {
			vw.RemoveObserver(interest, mediator)
		}
		delete (vw.mediatorMap, mediatorName)
		mediator.OnRemove();
	}
	return mediator;
}

func (vw *view) NotifyObservers(note api.Notification) {
	observers := vw.observerMap[note.GetName()]
	if (observers != nil) {
		for _, obs := range observers {
			obs.NotifyObserver(note)
		}
	}
}

func (vw *view) RegisterObserver(notificationName string, observer api.Observer) {
	obsArr, ok := vw.observerMap[notificationName]
	if ok {
		obsArr = append(obsArr, observer)
		vw.observerMap[notificationName] = obsArr
	} else {
		vw.observerMap[notificationName] = []api.Observer{observer}
	}
}

func (vw *view) RemoveObserver(notificationName string, notifyContext interface{}) {
	observers := vw.observerMap[notificationName];

	if (observers != nil) {
		for inx, obs := range observers {
			if (obs.CompareNotifyContext(notifyContext)) {
				if (len(observers) == 1) {
					delete(vw.observerMap, notificationName);
				} else {
					//remove index
					observers[inx] = observers[len(observers)-1]
					obs := observers[:len(observers)-1]
					vw.observerMap[notificationName] = obs	
				}
			}
		}
	}
}
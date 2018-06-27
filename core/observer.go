package core

import (
	"mvc/api"
)

type observer struct {
	context interface{};
	notify api.NotificationFunction
}

func NewObserver(notify api.NotificationFunction, context interface{}) *observer {
	obs := &observer{}
	obs.SetNotifyContext(context)
	obs.SetNotifyMethod(notify)
	return obs
}

func (obs *observer) CompareNotifyContext(object interface{}) bool {
	return obs.context == object;
}

func (obs *observer) NotifyObserver(notification api.Notification) {
	obs.notify(notification);
}

func (obs *observer) SetNotifyContext(notifyContext interface{}) {
	obs.context = notifyContext;
}

func (obs *observer) SetNotifyMethod(notifyMethod api.NotificationFunction) {
	obs.notify = notifyMethod;
}

package core

import (
	"github.com/mandeepbrar/mvc/interfaces"
)

type observer struct {
	context interface{}
	notify  interfaces.NotificationFunction
}

func NewObserver(notify interfaces.NotificationFunction, context interface{}) *observer {
	obs := &observer{}
	obs.SetNotifyContext(context)
	obs.SetNotifyMethod(notify)
	return obs
}

func (obs *observer) CompareNotifyContext(object interface{}) bool {
	return obs.context == object
}

func (obs *observer) NotifyObserver(notification interfaces.Notification) {
	obs.notify(notification)
}

func (obs *observer) SetNotifyContext(notifyContext interface{}) {
	obs.context = notifyContext
}

func (obs *observer) SetNotifyMethod(notifyMethod interfaces.NotificationFunction) {
	obs.notify = notifyMethod
}

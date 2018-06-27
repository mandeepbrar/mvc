package interfaces

type Observer interface {
	SetNotifyMethod(notifyMethod NotificationFunction)
	SetNotifyContext(notifyContext interface{})
	NotifyObserver(notification Notification)
	CompareNotifyContext(object interface{}) bool
}

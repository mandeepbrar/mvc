package interfaces

type Mediator interface {
	Notifier
	GetMediatorName() string
	GetViewComponent() interface{}
	SetViewComponent(viewComponent interface{})
	ListNotificationInterests() []string
	HandleNotification(notification Notification)
	OnRegister()
	OnRemove()
}

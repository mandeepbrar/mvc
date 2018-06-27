package interfaces

type Command interface {
	Notifier
	Execute(notification Notification)
}

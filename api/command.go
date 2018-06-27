package api

type Command interface{
	Notifier
	Execute(notification Notification)
}
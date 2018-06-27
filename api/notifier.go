package api

type Notifier interface {
	SendNotification(name string, body interface{}, typeName string)
	InitializeNotifier(key string)
}
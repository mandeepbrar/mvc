package interfaces

type Controller interface {
	ExecuteCommand(Notification)
	RegisterCommand(notificationName string, commandClassRef Command)
	HasCommand(notificationName string) bool
	RemoveCommand(notificationName string)
}

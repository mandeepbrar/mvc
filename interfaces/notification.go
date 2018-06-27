package interfaces

type NotificationFunction func(Notification)

type Notification interface {
	GetName() string
	GetBody() interface{}
	GetType() string
}

package api

type NotificationFunction func(Notification)

type Notification interface{
	GetName() string
	GetBody() interface{}
	GetType() string
}
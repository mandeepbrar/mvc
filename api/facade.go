package api

type Facade interface{
	Notifier
	RegisterCommand( notificationName string, commandClassRef Command)
	RemoveCommand( notificationName string)
	HasCommand( notificationName string ) bool
	RegisterProxy( proxy Proxy)
	RetrieveProxy(proxyName string) Proxy
	RemoveProxy(proxyName string) Proxy
	HasProxy( proxyName string) bool
	RegisterMediator(mediator Mediator)
	RetrieveMediator(mediatorName string) Mediator
	RemoveMediator(mediatorName string ) Mediator
	HasMediator(mediatorName string ) bool
	NotifyObservers(notification Notification)
}
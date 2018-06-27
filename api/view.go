package api

type View interface{
	RegisterObserver(notificationName string, observer Observer)
	RemoveObserver(notificationName string, notifyContext interface{})
	NotifyObservers(notification Notification )
	RegisterMediator(mediator Mediator)
	RetrieveMediator(mediatorName string) Mediator
	RemoveMediator(mediatorName string ) Mediator
	HasMediator(mediatorName string ) bool
}
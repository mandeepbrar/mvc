package patterns

import (
	"mvc/api"
)

type Mediator struct {
	*Notifier
	mediatorName string
	viewComponent interface{}
}

func NewMediator(mediatorName string, viewComponent interface{}) *Mediator {
	return &Mediator{NewNotifier(), mediatorName, viewComponent}
}

func (mediator *Mediator) GetMediatorName() string {
	return mediator.mediatorName
}

func (mediator *Mediator)GetViewComponent() interface{} {
	return mediator.viewComponent
}

func (mediator *Mediator)SetViewComponent(viewComponent interface{}) {
	mediator.viewComponent = viewComponent
}

func (mediator *Mediator)ListNotificationInterests() []string {
	return []string{}
}

func (mediator *Mediator)HandleNotification(notification api.Notification) {
}

func (mediator *Mediator)OnRegister() {

}

func (mediator *Mediator) onRemove() {

}

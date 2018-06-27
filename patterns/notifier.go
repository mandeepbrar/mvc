package patterns

type Notifier struct {
	multitonKey string
	facade      *facade
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (not *Notifier) GetFacade() *facade {
	if not.multitonKey == "" {
		panic("Notifier not initialized")
	}
	if not.facade == nil {
		not.facade = GetFacadeInstance(not.multitonKey)
	}
	return not.facade
}

func (not *Notifier) SendNotification(name string, body interface{}, typeName string) {
	fac := not.facade
	if fac == nil {
		fac = not.GetFacade()
	}
	if fac != nil {
		fac.SendNotification(name, body, typeName)
	}
}

func (not *Notifier) InitializeNotifier(key string) {
	not.multitonKey = key
}

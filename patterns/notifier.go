package patterns

type Notifier struct {
	multitonKey string
	fac *facade
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (not *Notifier) getFacade() *facade {
	if (not.multitonKey == "") {
		panic("Notifier not initialized");
	}
	if(not.fac == nil)	 {
		not.fac = GetFacadeInstance(not.multitonKey)
	}
	return not.fac
}

func (not *Notifier) SendNotification(name string, body interface{}, typeName string) {
	fac := not.fac
	if(fac == nil) {
		fac = not.getFacade()
	}
	if fac!=nil {
		fac.SendNotification(name, body, typeName)
	}
}

func (not *Notifier) InitializeNotifier(key string) {
	not.multitonKey = key;
}
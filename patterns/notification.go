package patterns

type notification struct {
	Name string
	Body interface{}
	Type string
}

func (not *notification) GetName() string {
	return not.Name
}

func (not *notification) GetBody() interface{} {
	return not.Body
}

func (not *notification) GetType() string {
	return not.Type
}
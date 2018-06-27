package api 

type Proxy interface {
	Notifier
	GetProxyName() string
	SetData(data interface{})
	GetData() interface{}
	OnRegister()
	OnRemove()
}
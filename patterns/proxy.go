package patterns

type Proxy struct {
	*Notifier
	name string
	data interface{}
}

func NewProxy(name string, data interface{}) *Proxy {
	return &Proxy{NewNotifier(), name, data}
}

func(proxy *Proxy) GetProxyName() string {
	return proxy.name
}

func(proxy *Proxy)SetData(data interface{}) {
	proxy.data = data
}

func(proxy *Proxy)GetData() interface{} {
	return proxy.data
}

func(proxy *Proxy)OnRegister() {

}

func(proxy *Proxy)OnRemove() {

}
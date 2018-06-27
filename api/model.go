package api

type Model interface {
	RegisterProxy(Proxy)
	RemoveProxy(proxyName string) Proxy
	RetrieveProxy(proxyName string) Proxy
	HasProxy( proxyName string) bool
}
package core

import (
	"mvc/api"
    "sync"
)

type model struct {
	proxyMap map[string]api.Proxy
	multitonKey string
}

var (
	mdlonce sync.Once
	mdlInstanceMap map[string]*model
)

func newModel(key string) *model {
	mdl := &model{make(map[string]api.Proxy), key}
	mdlInstanceMap[key] = mdl;
	mdl.InitializeModel();
	return mdl
}

func GetModelInstance(key string) *model {
    mdlonce.Do(func() {
        mdlInstanceMap = make(map[string]*model)
    })
	ins, ok := mdlInstanceMap[key]
	if !ok {
		return newModel(key)
	} else {
		return ins
	}
}

func (mdl *model) InitializeModel() {
}

func (mdl *model) RegisterProxy(proxy api.Proxy){
	proxy.InitializeNotifier(mdl.multitonKey)
	mdl.proxyMap[proxy.GetProxyName()] = proxy
	proxy.OnRegister()
}

func (mdl *model) RetrieveProxy(key string) api.Proxy {
	return mdl.proxyMap[key]
}
	

func (mdl *model) RemoveProxy(key string) api.Proxy {
	pxy, ok := mdl.proxyMap[key]
	if ok {
		delete(mdl.proxyMap, key)
		pxy.OnRemove()
	}
	return pxy
}
	

func RemoveModel(key string) {
	delete(mdlInstanceMap, key)
}
	

func (mdl *model) HasProxy(proxyName string) bool {
	_, ok := mdl.proxyMap[proxyName];
	return ok
}

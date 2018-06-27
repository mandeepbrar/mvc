package core

import (
	"sync"

	"github.com/mandeepbrar/mvc/interfaces"
)

type model struct {
	proxyMap    map[string]interfaces.Proxy
	multitonKey string
}

var (
	mdlonce        sync.Once
	mdlInstanceMap map[string]*model
)

func newModel(key string) *model {
	mdl := &model{make(map[string]interfaces.Proxy), key}
	mdlInstanceMap[key] = mdl
	mdl.InitializeModel()
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

func (mdl *model) RegisterProxy(proxy interfaces.Proxy) {
	proxy.InitializeNotifier(mdl.multitonKey)
	mdl.proxyMap[proxy.GetProxyName()] = proxy
	proxy.OnRegister()
}

func (mdl *model) RetrieveProxy(key string) interfaces.Proxy {
	return mdl.proxyMap[key]
}

func (mdl *model) RemoveProxy(key string) interfaces.Proxy {
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
	_, ok := mdl.proxyMap[proxyName]
	return ok
}

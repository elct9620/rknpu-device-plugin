package main

import "github.com/kubevirt/device-plugin-manager/pkg/dpm"

var _ dpm.ListerInterface = &Lister{}

type Lister struct {
}

func (l *Lister) GetResourceNamespace() string {
	return "rknpu"
}

func (l *Lister) Discover(chan dpm.PluginNameList) {
}

func (l *Lister) NewPlugin(name string) dpm.PluginInterface {
	return &Plugin{}
}

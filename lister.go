package main

import "github.com/kubevirt/device-plugin-manager/pkg/dpm"

var _ dpm.ListerInterface = &Lister{}

type Lister struct {
	ResUpdateChan chan dpm.PluginNameList
}

func NewLister() *Lister {
	return &Lister{
		ResUpdateChan: make(chan dpm.PluginNameList),
	}
}

func (l *Lister) GetResourceNamespace() string {
	return "rock-chips.com"
}

func (l *Lister) Discover(pluginListChan chan dpm.PluginNameList) {
	for {
		select {
		case detectedPlugins := <-l.ResUpdateChan:
			pluginListChan <- detectedPlugins
		case <-pluginListChan:
			return
		}
	}
}

func (l *Lister) NewPlugin(name string) dpm.PluginInterface {
	return &Plugin{}
}

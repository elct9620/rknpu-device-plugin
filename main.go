package main

import "github.com/kubevirt/device-plugin-manager/pkg/dpm"

func main() {
	lister := &Lister{}
	manager := dpm.NewManager(lister)

	manager.Run()
}

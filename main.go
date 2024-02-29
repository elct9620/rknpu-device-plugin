package main

import (
	"github.com/golang/glog"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	npuType, err := GetNpuType()
	if err != nil {
		glog.Fatalf("Failed to get NPU type: %v", err)
	} else {
		glog.Infof("Found NPU type: %s", npuType)
	}

	lister := &Lister{}
	manager := dpm.NewManager(lister)

	manager.Run()
}

package main

import (
	"github.com/golang/glog"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	defer glog.Flush()

	driverVersion, err := GetDriverVersion()
	if err != nil {
		glog.Errorf("failed to get driver version: %v", err)
		goto RUN
	}

	glog.Infof("NPU driver version: %s", driverVersion)

RUN:
	lister := &Lister{}
	manager := dpm.NewManager(lister)

	manager.Run()
}

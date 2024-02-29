package main

import (
	"os"

	"github.com/golang/glog"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	defer glog.Flush()

	lister := NewLister()
	manager := dpm.NewManager(lister)

	driverVersion, err := GetDriverVersion()
	if err != nil {
		glog.Errorf("failed to get driver version: %v", err)
		goto RUN
	}

	glog.Infof("NPU driver version: %s", driverVersion)

	go func() {
		if _, err := os.Stat(devicePath); os.IsNotExist(err) {
			glog.Errorf("NPU device not found: %s", devicePath)
			return
		}

		lister.ResUpdateChan <- dpm.PluginNameList{"rknpu"}
	}()

RUN:
	manager.Run()
}

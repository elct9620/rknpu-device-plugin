package main

import (
	"context"

	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

var _ dpm.PluginInterface = &Plugin{}

type Plugin struct {
}

func (p *Plugin) GetDevicePluginOptions(ctx context.Context, e *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	return &pluginapi.DevicePluginOptions{}, nil
}

func (p *Plugin) PreStartContainer(ctx context.Context, r *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	return &pluginapi.PreStartContainerResponse{}, nil
}

func (p *Plugin) ListAndWatch(e *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
	devices := []*pluginapi.Device{
		{
			ID:     devicePath,
			Health: pluginapi.Healthy,
		},
	}

	s.Send(&pluginapi.ListAndWatchResponse{Devices: devices})

	for {
		select {}
	}
}

func (p *Plugin) Allocate(ctx context.Context, r *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	var response pluginapi.AllocateResponse
	var car pluginapi.ContainerAllocateResponse
	var dev *pluginapi.DeviceSpec

	for _, req := range r.ContainerRequests {
		car = pluginapi.ContainerAllocateResponse{}

		for _, id := range req.DevicesIDs {
			dev = &pluginapi.DeviceSpec{
				ContainerPath: id,
				HostPath:      id,
				Permissions:   "rw",
			}

			car.Devices = append(car.Devices, dev)
		}

		response.ContainerResponses = append(response.ContainerResponses, &car)
	}

	return &response, nil
}

func (p *Plugin) GetPreferredAllocation(context.Context, *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error) {
	return &pluginapi.PreferredAllocationResponse{}, nil
}

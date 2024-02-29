package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	driverVersionPath = "/sys/module/rknpu/version"
	devicePath        = "/dev/dri/renderD129"
)

func GetDriverVersion() (string, error) {
	version, err := os.ReadFile(driverVersionPath)
	if err != nil {
		return "", fmt.Errorf("failed to read driver version: %v", err)
	}

	return strings.TrimSpace(string(version)), nil
}

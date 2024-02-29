package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	compatiblePath = "/proc/device-tree/compatible"
)

var (
	SupportedNpuTypes = []string{
		"rk3588",
	}
)

func GetNpuType() (string, error) {
	compatibleStr, err := os.ReadFile(compatiblePath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %v", compatiblePath, err)
	}

	for _, npuType := range SupportedNpuTypes {
		if strings.Contains(string(compatibleStr), npuType) {
			return npuType, nil
		}
	}

	return "", fmt.Errorf("no supported NPU type found in %s", compatiblePath)
}

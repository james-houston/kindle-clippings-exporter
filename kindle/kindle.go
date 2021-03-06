package kindle

import (
	"strings"

	usbdrivedetector "kindle_clipping_exporter/usbdrivedetector"
)

const kindleMountPointSuffix = "KINDLE"

func InitHandler() *Handler {
	handler := NewHandler()
	kindleMap := findAllKindles()
	for kindleMountPoint, ID := range kindleMap {
		handler.NewDevice(kindleMountPoint, ID)
	}
	return handler
}

func findAllKindles() map[string]string {
	usbDevices := usbdrivedetector.Detect()
	for deviceMountPoint := range usbDevices {
		if !strings.Contains(deviceMountPoint, kindleMountPointSuffix) {
			delete(usbDevices, deviceMountPoint)
		}
	}
	return usbDevices
}

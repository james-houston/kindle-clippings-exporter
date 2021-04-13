/*
	This file is a modified version of 1 file in Deepak Jois' golang usbdrivedetector
	Big thank you to him, you can view his original project here https://github.com/deepakjois/gousbdrivedetector
*/

package usbdrivedetector

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// Detect returns a list of file paths pointing to the root folder of
// kindle USB storage devices connected to the system.
func Detect() map[string]string {
	//drives := make(map[string]string)
	driveMap := make(map[string]string)
	dfPattern := regexp.MustCompile("^(\\/[^ ]+)[^%]+%[ ]+(.+)$")

	cmd := "df"
	out, err := exec.Command(cmd).Output()

	if err != nil {
		log.Printf("Error calling df: %s", err)
	}

	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		line := s.Text()
		if dfPattern.MatchString(line) {
			device := dfPattern.FindStringSubmatch(line)[1]
			rootPath := dfPattern.FindStringSubmatch(line)[2]

			if ok := isUSBStorage(device); ok {
				driveMap[rootPath] = getShortID(device)
			}
		}
	}

	return driveMap
}

func isUSBStorage(device string) bool {
	deviceVerifier := "ID_USB_DRIVER=usb-storage"
	cmd := "udevadm"
	args := []string{"info", "-q", "property", "-n", device}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		if device != "/dev/root" {
			// Don't log when checking /dev/root. Always an error unless running as root
			log.Printf("Error checking device %s: %s", device, err)
		}
		return false
	}

	if strings.Contains(string(out), deviceVerifier) {
		return true
	}

	return false
}

func getShortID(device string) string {
	idString := "ID_SERIAL_SHORT="
	cmd := "udevadm"
	args := []string{"info", "-q", "property", "-n", device}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		log.Printf("Error checking device %s: %s", device, err)
		return ""
	}

	parameters := strings.Split(string(out), "\n")
	for _, parameter := range parameters {
		if strings.Contains(parameter, idString) {
			return strings.TrimPrefix(parameter, idString)
		}
	}
	return ""
}

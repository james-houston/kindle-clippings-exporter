package kindle

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/james-houston/kindle-clippings-exporter/clippings"
)

const (
	// Default location for clippings test kindle (my kindle).
	// Able to add more later if there are other defaults.
	clippingsFile = "/documents/My Clippings.txt"
)

type Handler struct {
	NumDevices int
	Kindles    []Device
}

type Device struct {
	Path                string
	ClippingPath        string
	clippingsFileExists bool
	ID                  string
	allClippings        []*clippings.Clipping
}

func NewHandler() *Handler {
	return &Handler{}
}

// NewDevice creates a new device
func (h *Handler) NewDevice(mountPoint, ID string) {
	device := Device{
		Path: mountPoint,
		ID:   ID,
	}
	device.getClippingsPath()
	device.updateClippings()
	h.Kindles = append(h.Kindles, device)
	h.NumDevices++
	log.Printf("New device added to handler. %s. Current number of kindles connected: %d\n", device.ToString(), h.NumDevices)
}

// TODO: Eventually this should periodically check clippings
func (h *Handler) runListenerRoutine() {
	for {
		fmt.Println("listener routine running")
		time.Sleep(10 * time.Second)
	}
}

// ClippingPath checks for a "My Clippings.txt" file and returns the path to it
func (d *Device) getClippingsPath() {
	if _, err := os.Stat(d.Path + clippingsFile); err == nil {
		d.ClippingPath = d.Path + clippingsFile
		d.clippingsFileExists = true
		return
	} else {
		log.Printf("No clippings file for device %s: %v\n", d.ID, err)
	}
	d.ClippingPath = ""
	d.clippingsFileExists = false
}

func (d *Device) updateClippings() {
	if !d.clippingsFileExists {
		log.Printf("No clippings file for device %s. Skipping clippings update.\n", d.ID)
		return
	}
	d.allClippings = clippings.ParseClippingsFile(d.ClippingPath)
	log.Printf("First clipping: %v\n", d.allClippings[50])
}

func (d *Device) ToString() string {
	return fmt.Sprintf("Kindle ID: %s, mounted at %s, has clippings: %t, number of clippings: %d", d.ID, d.Path, d.clippingsFileExists, len(d.allClippings))
}

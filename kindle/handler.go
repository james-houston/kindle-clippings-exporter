package kindle

import (
	"fmt"
	"log"
	"os"

	"kindle_clipping_exporter/clippings"
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
	NewClippings        []*clippings.Clipping
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
	device.checkForNewClippings()

	h.Kindles = append(h.Kindles, device)
	h.NumDevices++
	log.Printf("New device added to handler. %s. Current number of kindles connected: %d\n", device.ToString(), h.NumDevices)
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
}

// Checks if any of the clippings are new to the disk
func (d *Device) checkForNewClippings() {
	for _, clipping := range d.allClippings {
		if clipping.IsNewClipping(d.ID) {
			log.Printf("New clipping in book %s\n", clipping.BookTitle)
			clipping.WriteToDisk(d.ID)
			if clipping.Type != clippings.TypeBookmark {
				d.NewClippings = append(d.NewClippings, clipping)
			}
		}
	}
}

//ToString returns a string with information on this device. Used for logging.
func (d *Device) ToString() string {
	return fmt.Sprintf("Kindle ID: %s, mounted at %s, has clippings: %t, number of clippings: %d", d.ID, d.Path, d.clippingsFileExists, len(d.allClippings))
}

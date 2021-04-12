package clippings

import (
	"log"
	"os"
	"path/filepath"
)

func (c *Clipping) WriteToDisk(kindleID string) {
	bookDir := filepath.Join(kindleID, c.BookTitle)
	if _, err := os.Stat(bookDir); os.IsNotExist(err) {
		if err := os.MkdirAll(bookDir, 0755); err != nil {
			log.Printf("Error creating directory for book %s: %v\n", c.BookTitle, err)
		}
	}

	clippingFile := filepath.Join(bookDir, c.timestampToString()+".txt")
	f, err := os.Create(clippingFile)
	if err != nil {
		log.Printf("Error creating new clipping file at %s: %v\n", clippingFile, err)
	}

	defer f.Close()

	_, err = f.WriteString(c.Body)

	if err != nil {
		log.Printf("Error writing clipping body to %s: %v", clippingFile, err)
	}

}

func (c *Clipping) IsNewClipping(kindleID string) bool {
	clippingFile := filepath.Join(".", kindleID, c.BookTitle, c.timestampToString()+".txt")
	if _, err := os.Stat(clippingFile); os.IsNotExist(err) {
		return true
	}

	// Otherwise, this clipping has already been created
	return false
}

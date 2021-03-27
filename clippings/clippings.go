package clippings

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Clipping struct {
	BookTitle string    // The book this clipping was taken from
	Author    string    // The author of the book
	Type      int       // The type of clipping (e.g. clipping/bookmark)
	Timestamp time.Time // Timestamp of the date the clipping was created
	Location  []int     // Location of the clipping. [0] = start point, [1] = end point. Bookmark is only [0]
	Body      string    // Actual text of the clipping
}

const clippingDelimiter = "=========="

const (
	typeBookmark = iota
	typeHighlight
	typeUnknown
)

func ParseClippingsFile(filePath string) []*Clipping {
	var allClippings []*Clipping
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error reading valid clippings file: %v \n", err)
		return nil
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == clippingDelimiter {
			// If the delimiter is the first line, we have no clippings. Exit.
			return nil
		}
		newClipping := &Clipping{}
		// first line is the title and author
		newClipping.parseTitleAndAuthor(scanner.Text())

		// second line is type, location, and date added
		scanner.Scan()
		newClipping.parseTypeAndLocationAndDate(scanner.Text())

		// third line is blank, ignore
		scanner.Scan()

		// fourth line is the text body
		scanner.Scan()
		newClipping.parseClippingBody(scanner.Text())

		// fifth line is the delimiter. Ignore.
		scanner.Scan()

		allClippings = append(allClippings, newClipping)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allClippings
}

func (c *Clipping) parseTitleAndAuthor(text string) {
	c.BookTitle = strings.Split(text, "(")[0]
	c.Author = strings.Split(text, "(")[1]
}

func (c *Clipping) parseTypeAndLocationAndDate(text string) {
	// Extract type from 3rd word of the line
	c.setType(strings.Split(text, " ")[2])

	// Extract location from 6th word line
	c.setLocation(strings.Split(text, " ")[5])

	// Extract timestamp by trimming prefix "... | Added on "
	tmp := strings.Split(text, "|")[1]
	tmp = strings.TrimPrefix(tmp, " Added on ")
	c.setTimestamp(tmp)
}

func (c *Clipping) setType(typeCompare string) {
	switch typeCompare {
	case "Bookmark":
		c.Type = typeBookmark
		break
	case "Hightlight":
		c.Type = typeHighlight
		break
	default:
		c.Type = typeUnknown
	}
}

// location is a string representation of a range in the form of two ints with a dash between them. e.g. "100-150"
func (c *Clipping) setLocation(location string) {
	splitLocation := strings.Split(location, "-")
	start, err := strconv.Atoi(splitLocation[0])
	if err != nil {
		log.Printf("Error getting start location for clipping: %v\n", err)
		start = 0
	}
	if len(splitLocation) == 1 {
		// Bookmark. Not a range of locaitons.
		c.Location = []int{start}
		return
	}
	end, err := strconv.Atoi(splitLocation[1])
	if err != nil {
		log.Printf("Error getting end location for a clipping: %v\n", err)
		end = 0
	}

	c.Location = []int{start, end}
}

func (c *Clipping) setTimestamp(timestamp string) {
	timestampFormat := "Monday, January 2, 2006 3:04:05 PM"
	if parsedTime, err := time.Parse(timestampFormat, timestamp); err == nil {
		c.Timestamp = parsedTime
	} else {
		log.Printf("Error parsing timestamp: %v", err)
	}
}

func (c *Clipping) timestampToString() string {
	return c.Timestamp.Format("2006-01-02_15:04:05")
}

func (c *Clipping) parseClippingBody(text string) {
	c.Body = text
}

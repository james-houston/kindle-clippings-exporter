package main

import (
	"log"

	"github.com/james-houston/kindle-clippings-exporter/kindle"
)

func main() {
	log.Println("Starting...")
	//kindleHandler := kindle.InitHandler()
	kindle.InitHandler()
}

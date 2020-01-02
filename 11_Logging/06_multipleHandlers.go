package main

import (
	"io"
	"log"
	"os"
)

func main() {

	file, e := os.OpenFile("06_multipleHandler.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Failed to open log file")
	}
	// You can set the logger output to a MultiWriter to write logs to multiple locations.
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)

	log.Println("message 1")

}

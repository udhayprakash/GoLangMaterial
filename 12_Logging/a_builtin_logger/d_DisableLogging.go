package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.Println("message 1")
	log.Println("message 2")

	// disable logging
	// SetOutput to set the output destination for the standard logger.
	log.SetOutput(ioutil.Discard)

	log.Println("message 3")
	log.Println("message 4")

	// re-enable logging
	log.SetOutput(os.Stderr)

	log.Println("message 5")
	log.Println("message 6")

}

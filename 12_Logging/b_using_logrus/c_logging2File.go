package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("c_logging.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Set logrus to write to the file
	log.SetOutput(file)

	// redirecting the logs the file
	log.SetOutput(file)
	log.Debug("This is an debug log")
	log.Info("This is an info log")
	log.Warning("This is a warning log")
	log.Error("This is an error log")

}

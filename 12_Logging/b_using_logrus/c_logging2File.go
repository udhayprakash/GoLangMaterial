package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	file, err := os.OpenFile("c_logging.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// redirecting the logs the file
	log.SetOutput(file)
	log.Debug("This is an debug log")
	log.Info("This is an info log")
	log.Warning("This is a warning log")
	log.Error("This is an error log")

}

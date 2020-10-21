package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	file, err := os.OpenFile("d_setLogFormat_as_json.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// redirecting the logs the file
	log.SetOutput(file)

	// default log level is INFO. To set the level
	log.SetLevel(log.WarnLevel)

	// formatting the log
	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("This is an debug log")
	log.Info("This is an info log")
	log.Warning("This is a warning log")
	log.Error("This is an error log")


	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

}

package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("d_setLogFormat_as_json.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()

	// Redirect logs to the file
	log.SetOutput(file)

	// Set log level to DebugLevel to see all logs
	log.SetLevel(log.DebugLevel)

	// Format logs as JSON
	log.SetFormatter(&log.JSONFormatter{})

	// Log messages
	log.Debug("This is a debug log")
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
	
	fmt.Println("end of file, wont run, as Fatal will be last line to execute")

}

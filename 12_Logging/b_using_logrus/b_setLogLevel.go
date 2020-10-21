package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// default log level is INFO. To set the level
	log.SetLevel(log.WarnLevel)

	log.Debug("This is an debug log")
	log.Info("This is an info log")
	log.Warning("This is a warning log")
	log.Error("This is an error log")

}

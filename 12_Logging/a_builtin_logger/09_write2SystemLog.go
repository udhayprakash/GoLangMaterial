package main

import (
	"log"
	"log/syslog"
)

// The syslog package provides a simple interface to the system log service.
func main() {

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
	if e == nil {
		log.SetOutput(logwriter)
	}

	log.Print("Hello System Log")
}

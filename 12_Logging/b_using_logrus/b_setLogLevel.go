package main

import (
	log "github.com/sirupsen/logrus"
)

// Unlike the standard log package, logrus supports log levels.
// logrus has seven log levels: Trace, Debug, Info, Warn, Error, Fatal, and Panic.
func main() {
	// default log level is INFO. To set the level
	log.SetLevel(log.DebugLevel) // log.WarnLevel

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}

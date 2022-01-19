package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	log.SetFormatter(Formatter)

	log.Info("This is info log")
	log.Warning("This is a warning")
	log.Error("Not fatal. An error. Won't stop execution")
	log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	log.Panic("Do not panic")
}

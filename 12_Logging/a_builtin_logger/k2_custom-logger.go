package main

import (
	"log"
	"os"
)

type AggregateLogger struct {
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *AggregateLogger) warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *AggregateLogger) info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *AggregateLogger) error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

func main() {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	agLog := AggregateLogger{
		warnLogger:  log.New(os.Stdout, "WARNING: ", flags),
		infoLogger:  log.New(os.Stdout, "INFO   : ", flags),
		errorLogger: log.New(os.Stdout, "ERROR  : ", flags),
	}
	agLog.warn("This is warning log")
	agLog.info("This is information log")
	agLog.error("This is error log")
}

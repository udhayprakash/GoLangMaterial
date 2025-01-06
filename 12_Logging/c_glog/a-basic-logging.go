package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// Parse command-line flags (required by glog)
	flag.Parse()

	// Log messages at different levels
	glog.Info("This is an info log")
	glog.Warning("This is a warning log")
	glog.Error("This is an error log")

	// Flush logs to ensure all buffered logs are written
	glog.Flush()
}

// $ go run a-basic-logging.go
// $ go run a-basic-logging.go  -logtostderr=true

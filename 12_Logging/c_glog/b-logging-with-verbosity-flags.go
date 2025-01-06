package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Log messages with different verbosity levels
	glog.V(1).Info("This is a V1 log")
	glog.V(2).Info("This is a V2 log")
	glog.V(3).Info("This is a V3 log")

	// Flush logs
	glog.Flush()
}

// $ go run b-logging-with-verbosity-flags.go -logtostderr=true -v=2

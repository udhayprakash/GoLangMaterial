package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Conditional logging
	debugMode := true
	if debugMode {
		glog.Info("Debug mode is enabled")
	} else {
		glog.Warning("Debug mode is disabled")
	}

	// Flush logs
	glog.Flush()
}

// $ go run d-conditional-logging.go  -logtostderr=true
// # command-line-arguments
// ./d-conditional-logging.go:15:7: undefined: glog.InfoIf
// ./d-conditional-logging.go:16:7: undefined: glog.WarningIf

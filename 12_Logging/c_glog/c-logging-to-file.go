package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/golang/glog"
)

func init() {
	// Call the shell script to clean up old log files
	cmd := exec.Command("./cleanup_logs.sh")
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to run cleanup script: %v\n", err)
	}
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Log messages
	glog.Info("This log will be written to a file")
	glog.Warning("This is a warning log in a file")

	// Flush logs
	glog.Flush()
}

/*

	$ go run c-logging-to-file.go


By default logs to

	Linux/macOS: /tmp
	Windows: C:\Users\<YourUser>\AppData\Local\Temp

	$ go run c-logging-to-file.go -log_dir=./logs


mkdir logs
	$ go run c-logging-to-file.go -log_dir=./logs



Two sets of files created
	1) one file , each, for each log level, defined in code
	2) one file for every execution
		rotated files, with timestamp and ProcessID

$ go run c-logging-to-file.go -log_dir=./logs -log_file_max_count=3
	// there is small bug

$ chmod +x cleanup_logs.sh
$ go run c-logging-to-file.go -log_dir=./logs

*/

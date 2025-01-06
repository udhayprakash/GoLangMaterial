package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// CustomJSONFormatter is a custom formatter to ensure "time" is the first key in JSON output.
type CustomJSONFormatter struct {
	log.JSONFormatter
}

// Format implements the logrus.Formatter interface.
func (f *CustomJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Create a map to hold the log data
	data := make(log.Fields, len(entry.Data)+3) // +3 for time, level, and msg

	// Add the time first
	data["time"] = entry.Time.Format(time.RFC3339)

	// Add the level and message
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message

	// Add the rest of the fields
	for k, v := range entry.Data {
		data[k] = v
	}

	// Convert the map to JSON
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON: %v", err)
	}

	return b.Bytes(), nil
}

func main() {
	file, err := os.OpenFile("e_setLogFormat_as_json.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer file.Close()

	// Redirect logs to the file
	log.SetOutput(file)

	// Set log level to DebugLevel to see all logs
	log.SetLevel(log.DebugLevel)

	// Use the custom JSON formatter
	log.SetFormatter(&CustomJSONFormatter{})

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

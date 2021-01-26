package main

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// // Context interface
// type Context interface {
// 	Done() <-chan struct{}
// 	Deadline() (deadline time.Time, ok bool)
// 	Err() error
// 	Value(key interface{}) interface{}
// }

func main() {

	type keyval string

	// Generating Parent-Context
	ctx := context.Background()

	// Here, context is created with values
	vctx := context.WithValue(ctx, keyval("request-id"), keyval("123"))

	// Preconfigure Logrus Logger
	requestLogger := log.WithFields(
		log.Fields{
			"request-id": vctx.Value(keyval("request-id")),
		})

	// Log without specifying further values
	// these are preconfigured in the logger
	requestLogger.Info("Some important message")
	requestLogger.Info("Next important message")

}

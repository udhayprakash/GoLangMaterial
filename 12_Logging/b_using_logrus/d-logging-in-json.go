package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Println("Hello world1")

	log.SetFormatter(&log.JSONFormatter{})
	log.Println("Hello world2")

	// adding fields to log
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")

}

package main

import (
	"errors"
	"log"
)

var exampleError = errors.New("Example Error, DoSomething has gone wrong!")

func DoSomeThing() error {
	return exampleError
}

func main() {

	err := DoSomeThing()

	if errors.Is(err, exampleError) {
		log.Println("The error was our example error")
	} else if err != nil {
		log.Println("The error was something else")
	} else {
		log.Println("There was no error")
	}
}

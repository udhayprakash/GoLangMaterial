package main

import (
	"fmt"
	"log"
)

func cleanUp() error {
	return fmt.Errorf("**ERROR: cleanUp() error%s", "Udhay")
}

func getMessageBug() (s string, err error) {
	s = "Ok"

	fmt.Println(&err, err)
	defer func() {
		err = cleanUp()
		s = "This too is buggy"
		fmt.Println(&err, err)
	}()

	return
}

func main() {

	msg, err := getMessageBug()
	if err != nil {
		log.Println(err)
	}
	log.Println(msg)
}

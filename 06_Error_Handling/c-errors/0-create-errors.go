package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	var er error
	fmt.Printf("%%s: %s\n%%v: %v\n%%#v: %#v", er, er, er)

	err := errors.New("Zero Division Error")
	fmt.Println("\nError is :", err)

	fmt.Printf(`
		Error Details
		value : %[1]v
		Type  : %[1]T
		type  : %[2]v
	`, err, reflect.TypeOf(err).Kind())

	timeErr := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("\nError is :", timeErr)

	fmt.Printf(`
		Error Details
		value : %[1]v
		Type  : %[1]T
		type  : %[2]v
	`, timeErr, reflect.TypeOf(timeErr).Kind())
}

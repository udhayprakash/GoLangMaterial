package main

import (
	"errors"
	"fmt"
	"os"
)

type RequestError struct {
	StatusCode int

	Err error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func main() {
	err := doRequest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("success!")
}

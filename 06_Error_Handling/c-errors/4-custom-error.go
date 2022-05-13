package main

import (
	"errors"
	"fmt"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func doRequest() (string, error) {
	return "", &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func main() {
	msg, err := doRequest()
	if err != nil {
		fmt.Println("Error is -", err)
		return
	}
	fmt.Println("Msg:", msg)

}

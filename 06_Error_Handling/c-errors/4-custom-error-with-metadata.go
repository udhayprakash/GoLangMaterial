package main

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	StatusCode int
	Err        error
}

func (r *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %v", r.StatusCode, r.Err)
}

func makeHTTPRequest() (string, error) {
	return "", &HTTPError{
		StatusCode: 404,
		Err:        errors.New("not found"),
	}
}

func main() {
	msg, err := makeHTTPRequest()
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	fmt.Println("Response:", msg)
}

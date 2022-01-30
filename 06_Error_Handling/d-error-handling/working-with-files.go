package main

import (
	"fmt"
	"os"
	"time"
)

type ErrorWithTime struct {
	err error     // the wrapped error
	t   time.Time // the time when the error happened
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("%v @ %s", e.err, e.t)
}

func wrap(err error) *ErrorWithTime {
	return &ErrorWithTime{err, time.Now()}
}

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, wrap(err)
	}

	return f, nil
}

func main() {
	file, err := openFile("notes.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
}

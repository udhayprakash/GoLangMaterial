package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// Error Handling
type SyntaxError struct {
	Line int
	Col  int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}

type InternalError struct {
	Path string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("parse %v:internal error", e.Path)
}

func main() {
	if err := Foo(); err != nil {
		switch e := err.(type) {
		case *SyntaxError:
			// Do something interesting woth e.Line and e.Col
		case *InternalError:
			// Abort and file an issue
		default:
			log.Println(e)
		}
	}
}

func Foo() error {
	// simple string-based error
	err1 := errors.New("math: square root of negative number")
	return err1

	// with formatting 
	err2 := fmt.Errorf("math: square root of negative number %g", -12)
	return err2
	
}

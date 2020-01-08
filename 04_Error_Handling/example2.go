package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Zero Division Error")
	fmt.Println("Error is :", err)
}

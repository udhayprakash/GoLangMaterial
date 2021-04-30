package main

import (
	"fmt"
)

type String string

func (s String) Required() bool {

	return s != ""
}

func main() {
	angle := "udhay"
	fmt.Println("angle:", angle)

	fmt.Println(String("abcd").Required())
}

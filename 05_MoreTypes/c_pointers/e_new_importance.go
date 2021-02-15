package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person
	p2 := new(Person)

	p1.name = "peeOne"
	p2.name = "peeTwo"

	fmt.Println("p1=", p1, reflect.TypeOf(p1).Kind()) // struct
	fmt.Println("p2=", p2, reflect.TypeOf(p2).Kind()) // pointer
}

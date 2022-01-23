package main

import (
	"fmt"
)

/*
Purpose: Interfaces
	- are named collections of method signatures.
	- Polymorphism can be achieved using interfaces.

SYNTAX:
	type InterfaceName interface{
		MethodName(argument argumentType) returnType
	}


To implement an interface in Go, we just need to
implement all the methods in the interface.
*/
// 1) Interfaces
type Bird interface {
	Fly()
}

// 2) Structs & Methods
type Eagle struct{}

func (e Eagle) Fly() {
	fmt.Println("Eagle flies over the cloud")
}

type Crow struct{}

func (c Crow) Fly() {
	fmt.Println("Crow flies at normal height")
}

type Hen struct{}

func (h Hen) Fly() {
	fmt.Println("Hen can not fly")
}

// 3) Functions with interface type

func canYouFly(b Bird) {
	b.Fly()
}

func main() {
	canYouFly(Eagle{})
	canYouFly(Crow{})
	canYouFly(Hen{})
}

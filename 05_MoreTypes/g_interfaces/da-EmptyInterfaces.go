package main

import "fmt"

// -------------
type Eagle1 struct{}

func (e Eagle1) Fly() {
	fmt.Println("Eagle flies over the cloud")
}

// ------------
type Crow1 struct{}

func (c Crow1) Fly() {
	fmt.Println("Crow flies at normal height")
}

// ---------
type Hen1 struct{}

func (h Hen1) Fly() {
	fmt.Println("Hen can not fly")
}

// ===========================
// Functions with interface type
// Empty interface can just be used only

func canYouFly1(a interface{}) {
	fmt.Println("a =", a)
	// a.Fly()  // a.Fly undefined (type interface {} is interface with no methods)
}

func main() {
	canYouFly1(Eagle1{})
	canYouFly1(Crow1{})
	canYouFly1(Hen1{})
}

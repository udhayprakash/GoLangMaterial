package main

import (
	"fmt"
)

type math interface {
	sum() float64
	substr() float64
}

type item1 struct {
	num1, num2 float64
}

type item2 struct {
	num1, num2 float64
}

// Method
func (a item1) sum() float64 {
	return a.num1 + a.num2
}

// Method
func (a item1) substr() float64 {
	return a.num1 - a.num2
}

// Method
func (b item2) sum() float64 {
	return b.num1 + b.num2
}

// Method
func (b item2) substr() float64 {
	return b.num1 - b.num2
}

func result(m math) {
	fmt.Println("Sum   : ", m.sum())
	fmt.Println("Substr: ", m.substr())
}

func main() {
	a := item1{num1: 10, num2: 5}
	b := item2{num1: 20, num2: 10}

	result(a)
	result(b)
}

// Sum   :  15
// Substr:  5
// Sum   :  30
// Substr:  10

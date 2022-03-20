package main

import "fmt"

type item1 struct {
	num1, num2 float64
}

func (a item1) sum() float64 {
	return a.num1 + a.num2
}
func (a item1) substr() float64 {
	return a.num1 - a.num2
}

type item2 struct {
	num1, num2 float64
}

func (b item2) sum() float64 {
	return b.num1 + b.num2
}
func (b item2) substr() float64 {
	return b.num1 - b.num2
}

type math interface {
	sum() float64
	substr() float64
}

func result(m math) {
	fmt.Println("Sum   : ", m.sum())
	fmt.Println("Substr: ", m.substr())
}

func main() {
	a := item1{num1: 10, num2: 5}
	result(a)

	b := item2{num1: 20, num2: 10}
	result(b)
}

//Sum   :  15
//Substr:  5
//Sum   :  30
//Substr:  10

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ... A composite literal allows you to
	// assign a value directly to an array,
	// slice, or map
	a1 := [...]int{1, 2, 3}
	a2 := [...]float64{1.0, 2.0, 3.0}
	a3 := [...]string{"a", "b", "c"}
	a4 := [...]rune{'a', 'b', 'c'}

	fmt.Println("a1=", a1, reflect.TypeOf(a1).Kind())
	fmt.Println("a2=", a2, reflect.TypeOf(a2).Kind())
	fmt.Println("a3=", a3, reflect.TypeOf(a3).Kind())
	fmt.Println("a4=", a4, reflect.TypeOf(a4).Kind())

}

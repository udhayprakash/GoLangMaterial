package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{11, 22, 33} // [11 22 33] - [3]int - array
	fmt.Printf("%v - %[1]T - %v\n", a1, reflect.TypeOf(a1).Kind())

	a2 := [...]int{11, 22, 33} // [11 22 33] - [3]int - array
	fmt.Printf("%v - %[1]T - %v\n", a2, reflect.TypeOf(a2).Kind())

	a3 := []int{11, 22, 33} // [11 22 33] - []int - slice
	fmt.Printf("%v - %[1]T - %v\n", a3, reflect.TypeOf(a3).Kind())

	a4 := make([]int, 3) // [0 0 0] - []int - slice
	fmt.Printf("%v - %[1]T - %v\n", a4, reflect.TypeOf(a4).Kind())

	fmt.Println("a1 == a2 :", a1 == a2) // true
	// a1 == a3 // Can't compare array & slice

}
 
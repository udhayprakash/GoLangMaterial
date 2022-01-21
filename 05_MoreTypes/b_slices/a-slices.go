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

	s1 := []int{11, 22, 33} // [11 22 33] - []int - slice
	fmt.Printf("%v - %[1]T - %v\n", s1, reflect.TypeOf(s1).Kind())

	s2 := make([]int, 3) // [0 0 0] - []int - slice
	fmt.Printf("%v - %[1]T   - %v\n", s2, reflect.TypeOf(s2).Kind())

	fmt.Println("a1 == a2 :", a1 == a2) // true
	// a1 == a3 // Can't compare array & slice

	// Also, slice can be created from an array
	var s3 []int = a1[1:3]
	fmt.Printf("%v - %[1]T   - %v\n", s3, reflect.TypeOf(s3).Kind())

}

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
	fmt.Printf("%v - %[1]T  - %v\n", s1, reflect.TypeOf(s1).Kind())

	s2 := make([]int, 3) // [0 0 0] - []int - slice
	fmt.Printf("%v    - %[1]T  - %v\n", s2, reflect.TypeOf(s2).Kind())

	// Also, slice can be created from an array
	var s3 []int = a1[1:3]
	fmt.Printf("%v    - %[1]T  - %v\n", s3, reflect.TypeOf(s3).Kind())
	fmt.Println()

	// comparision
	fmt.Println("a1 == a2  :", a1 == a2) // true
	// a1 == s3 // Can't compare array & slice
	// s1 == s2 // slice can only be compared to nil

	fmt.Println("s1 == nil :", s1 == nil) // false
	fmt.Println("len(s1)   :", len(s1))

	fmt.Println("s2 == nil :", s2 == nil) // false
	fmt.Println("len(s2)   :", len(s2))

	s4 := make([]int, 0)
	fmt.Printf("%v    - %[1]T  - %v\n", s4, reflect.TypeOf(s4).Kind())
	fmt.Printf("len(s4)=%d \ts4 == nil =%t\n", len(s4), s4 == nil) // 0 false

	var s5 []int // nil slice
	fmt.Printf("%v    - %[1]T  - %v\n", s5, reflect.TypeOf(s5).Kind())
	fmt.Printf("len(s5)=%d \ts5 == nil =%t\n", len(s5), s5 == nil) // 0 true
}

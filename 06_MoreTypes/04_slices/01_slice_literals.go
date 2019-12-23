package main

import (
	"fmt"
	"reflect"
)

/*
array literal syntax: [3]bool{true, true, false}
array literal syntax: [...]bool{true, true, false}
slice literal syntax: []bool{true, true, false}

*/
func main() {
	array1 := [3]int{2, 3, 5}
	fmt.Println(reflect.TypeOf(array1).Kind(), array1)

	array2 := [...]bool{false, true, true}
	fmt.Println(reflect.TypeOf(array2).Kind(), array2)

	mySlice1 := []bool{false, true, true}
	fmt.Println(reflect.TypeOf(mySlice1).Kind(), mySlice1)

	mySlice2 := make([]string, 3)
	fmt.Println(reflect.TypeOf(mySlice2).Kind(), mySlice2) // [  ]

	mySlice3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true}, // comma is essential, here
	}
	fmt.Println(reflect.TypeOf(mySlice3).Kind(), mySlice3)
}

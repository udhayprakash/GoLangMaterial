package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [3]int{2, 3, 5}
	fmt.Println(reflect.TypeOf(array1).Kind(), array1)

	var array3 = []bool{false, true, true}
	fmt.Println(reflect.TypeOf(array3).Kind(), array3)

	array2 := [...]bool{false, true, true}
	fmt.Println(reflect.TypeOf(array2).Kind(), array2)

	mySlice1 := []bool{false, true, true}
	fmt.Println(reflect.TypeOf(mySlice1).Kind(), mySlice1)

	mySlice2 := make([]string, 3)
	fmt.Println(reflect.TypeOf(mySlice2).Kind(), mySlice2) // [  ]
	fmt.Printf("myslice2 - capacity:%d 	length:%d\n", cap(mySlice2), len(mySlice2))

	mySlice3 := make([]string, 3, 10)
	fmt.Println(reflect.TypeOf(mySlice3).Kind(), mySlice3) // [  ]
	fmt.Printf("mySlice3 - capacity:%d 	length:%d\n", cap(mySlice3), len(mySlice3))

	mySlice4 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true}, // comma is essential, here
	}
	fmt.Println(reflect.TypeOf(mySlice4).Kind(), mySlice4)

	var s5 []int
	fmt.Println(len(s5), len(s5) == 0, s5 == nil)
	s5 = nil
	fmt.Println(len(s5), len(s5) == 0, s5 == nil)
	s5 = []int(nil)
	fmt.Println(len(s5), len(s5) == 0, s5 == nil)
	s5 = []int{}
	fmt.Println(len(s5), len(s5) == 0, s5 != nil)
}
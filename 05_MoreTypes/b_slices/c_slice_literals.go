package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [3]int{2, 3, 5} // array
	fmt.Println(reflect.TypeOf(array1).Kind(), array1)

	array2 := [...]bool{false, true, true} // array
	fmt.Println(reflect.TypeOf(array2).Kind(), array2)

	mySlice1 := []bool{false, true, true} //slice
	fmt.Println(reflect.TypeOf(mySlice1).Kind(), mySlice1)

	mySlice2 := make([]int, 3)                             //slice
	fmt.Println(reflect.TypeOf(mySlice2).Kind(), mySlice2) // slice [0 0 0]
	fmt.Printf("myslice2 - capacity:%d 	length:%d\n", cap(mySlice2), len(mySlice2))

	mySlice3 := make([]string, 3, 10)                      // slice
	fmt.Println(reflect.TypeOf(mySlice3).Kind(), mySlice3) // slice [  ]
	fmt.Printf("mySlice3 - capacity:%d 	length:%d\n", cap(mySlice3), len(mySlice3))

	mySlice4 := new([30]int)[0:10]
	fmt.Println(reflect.TypeOf(mySlice4).Kind(), mySlice4) // slice [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("mySlice4 - capacity:%d 	length:%d\n", cap(mySlice4), len(mySlice4))
	fmt.Println()

	mySlice5 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true}, // comma is essential, here
	}
	fmt.Println(reflect.TypeOf(mySlice5).Kind(), mySlice5)

	var s5 []int
	fmt.Println(len(s5), len(s5) == 0, s5 == nil)
	s5 = nil
	fmt.Println(len(s5), len(s5) == 0, s5 == nil)
	s5 = []int(nil)
	fmt.Println(len(s5), len(s5) == 0, s5 == nil)
	s5 = []int{}
	fmt.Println(len(s5), len(s5) == 0, s5 != nil)

}

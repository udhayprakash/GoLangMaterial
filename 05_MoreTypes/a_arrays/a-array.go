package main

import (
	"fmt"
	"reflect"
)

/*
Array - data structure  to store a collection of data of the same type
Declaration
	var array_name [length]Type
	or
	var array_name[length]Type{item1, item2, item3, ...itemN}

Short declaration
	array_name:= [length]Type{item1, item2, item3,...itemN}

*/

func main() {
	// Declaration of Arrays
	var emptyArray1 [3]int
	fmt.Println("emptyArray1 = ", emptyArray1) // [0 0 0]

	var emptyArray2 [4]bool
	fmt.Println("emptyArray2 = ", emptyArray2) // [false false false false]

	var emptyArray3 [2]string
	fmt.Println("emptyArray3 =", emptyArray3) // ["" ""]
	fmt.Println()

	// Declaration & initialization
	var myArray1 [3]int = [3]int{11, 22, 33}
	fmt.Println("myArray1    = ", myArray1)

	var myArray2 = [3]int{11, 22, 33}
	fmt.Println("myArray2    = ", myArray2)

	myArray3 := [3]int{11, 22, 33}
	fmt.Println("myArray3    = ", myArray3)
	fmt.Println()

	// Multi-line initialization
	greetings1 := [4]string{"Good Morning", "Good Afternoon", "Good Evening", "Good Night"}
	fmt.Println("greetings1    = ", greetings1)

	greetings2 := [4]string{
		"Good Morning",
		"Good Afternoon",
		"Good Evening",
		"Good Night", // last comma is needed
	}
	fmt.Println("greetings2    = ", greetings2)
	// [Good Morning Good Afternoon Good Evening Good Night]

	// Automatic array length declaration
	// when length is represented by ...
	// Go compiler will find the length on its own
	// ... can be used when defining array with initial value
	greetings3 := [...]string{
		"Good Morning",
		"Good Afternoon",
		"Good Evening",
		"Good Night", // last comma is needed
	}
	fmt.Println("greetings3    =", greetings3)

	fmt.Println("len(greetings1)  :", len(greetings1))
	fmt.Println("len(greetings2)  :", len(greetings2))
	fmt.Println("len(greetings3)  :", len(greetings3))

	fmt.Printf("\n\ngreetings1 - type=%T", greetings1)
	fmt.Printf("\ngreetings1 - type=%T", greetings2)
	fmt.Printf("\ngreetings1 - type=%T", greetings3)

	fmt.Println(reflect.TypeOf(greetings1))
}

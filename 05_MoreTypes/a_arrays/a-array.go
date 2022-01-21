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
	// Method 1 - Declaration of Arrays
	var emptyArray1 [3]int
	fmt.Println("emptyArray1 = ", emptyArray1) // [0 0 0]

	var emptyArray2 [4]bool
	fmt.Println("emptyArray2 = ", emptyArray2) // [false false false false]

	var emptyArray3 [2]string
	fmt.Println("emptyArray3 =", emptyArray3) // ["" ""]
	fmt.Println()

	// Method2 - Declaration & initialization
	var myArray1 [3]int = [3]int{11, 22, 33}
	fmt.Println("myArray1    = ", myArray1)

	var myArray2 = [3]int{11, 22, 33}
	fmt.Println("myArray2    = ", myArray2)

	myArray3 := [3]int{11, 22, 33}
	fmt.Println("myArray3    = ", myArray3) // [11 22 33]

	myArray4 := [3]int{11, 22}
	fmt.Println("myArray4    = ", myArray4) // [11 22 0]
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

	fmt.Println("\n", reflect.TypeOf(greetings1))

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

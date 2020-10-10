package main

import "fmt"

func main() {
	var emptyArray1[3]int
	fmt.Println("emptyArray1 :", emptyArray1) // [0 0 0]

	var emptyArray2[2]bool
	fmt.Println("emptyArray2 :", emptyArray2) // [false false]

	// Declaration & initialization
	var myArray3 [3]int = [3]int{11, 22, 33}
	fmt.Println("myArray3    :", myArray3) // [11 22 33]

	var myArray4 = [3]int{44, 66, 55}
	fmt.Println("myArray4    :", myArray4) // [44 66 55]

	myArray5  := [3]int{44, 66, 55}
	fmt.Println("myArray5    :", myArray5) // [44 66 55]

	// Multi-line initialization
	greetings := [4]string{
		"Good Morning",
		"Good Afternoon",
		"Good Evening",
		"Good Night",  // last comma is needed
	}
	fmt.Println("greetings    :", greetings) // [Good Morning Good Afternoon Good Evening Good Night]

	// Automatic array length declaration
	// when length is represented by ...
	// Go compiler will find the length on its own
	// ... can be used when defining array with initial value
	greetings1 := [...]string{
		"Good Morning",
		"Good Afternoon",
		"Good Evening",
		"Good Night",  // last comma is needed
	}
	fmt.Println("greetings1   :", greetings1) // [Good Morning Good Afternoon Good Evening Good Night]

	fmt.Println("len(greetings)  :", len(greetings))
	fmt.Println("len(greetings1) :", len(greetings1))
}

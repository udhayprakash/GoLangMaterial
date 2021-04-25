package main

import (
	"fmt"
)

// Maps are of reference type
func main() {
	num1 := 123
	num2 := num1
	fmt.Println("\nbefore copy - num1=", num1, "\t num2=", num2)

	num1 = 777
	fmt.Println("after copy - num1=", num1, "\t num2=", num2)

	name1 := "udhay"
	var name2 string = name1
	fmt.Printf("\nbefore copy - name1 = %s \t name2 = %s\n", name1, name2)

	name2 = "Prakash"
	fmt.Printf("after copy - name1 = %s \t name2 = %s\n", name1, name2)


	// maps are reference types
	var myMap1 = map[string]int{
		"one": 1, 
		"two": 2, 
		"three": 3,
	}

	myMap2 := myMap1 
	fmt.Printf("\nBefore copy - myMap1=%v \n\t\t myMap2=%v\n", myMap1, myMap2)

	myMap2["three"] = 333

	fmt.Printf("After copy - myMap1=%v \n\t\t myMap2=%v", myMap1, myMap2)

	// Better way to copy map 
	myMap3 := make(map[string]int)
	for key, value := range myMap1 {
		myMap3[key] = value
	}
	fmt.Printf("\nBefore copy - myMap1=%v \n\t\t myMap3=%v\n", myMap1, myMap3)

	myMap3["one"] = 111
	fmt.Printf("After copy - myMap1=%v \n\t\t myMap3=%v", myMap1, myMap3)

}

package main

import (
	"fmt"
)

func main() {
	// Method 1 - Declaring, and  initializing the map
	studentRank1 := map[string]int{}
	fmt.Println("\nstudentRank1      = ", studentRank1)
	fmt.Println("len(studentRank1) = ", len(studentRank1))

	// Method2 - Declaring, and  initializing the map
	studentRank := make(map[string]int)
	fmt.Println("\nstudentRank      = ", studentRank)
	fmt.Println("len(studentRank) = ", len(studentRank))

	// To add key-value pairs
	studentRank["Suresh"] = 1
	studentRank["Mahesh"] = 3
	studentRank["Ganesh"] = 2

	fmt.Println("\nstudentRank      = ", studentRank)
	fmt.Println("len(studentRank) = ", len(studentRank))

	// To overwrite all key-value pairs
	studentRank = map[string]int{"Apparao": 3, "subbarao": 1, "ramarao": 2}
	fmt.Println("\nstudentRank      = ", studentRank)
	fmt.Println("len(studentRank) = ", len(studentRank))

	// maps are indexed by keys, not position
	fmt.Println("\nstudentRank[\"subbarao\"]=", studentRank["subbarao"])
	fmt.Println("studentRank[\"Mahesh\"]=", studentRank["Mahesh"])

	// To check for present of a key
	val, isKeyPresent := studentRank["ramarao"]
	fmt.Printf("\nval=%v \t isKeyPresent=%v\n", val, isKeyPresent)

	val, isKeyPresent = studentRank["superman"]
	fmt.Printf("val=%v \t isKeyPresent=%v\n", val, isKeyPresent)

	val = studentRank["Apparao"]
	fmt.Println("val=", val)

	// blank identifier _ - used when we dont need a value
	_, isKeyPresent = studentRank["superman"]
	fmt.Println("isKeyPresent:", isKeyPresent) // false

	// To delete a key-value pair
	fmt.Println("\nBefore delete, studentRank=", studentRank)
	delete(studentRank, "Apparao")

	// deleting a key which doesnt exits, won't through an error
	delete(studentRank, "rajinikanth")
	fmt.Println("After delete, studentRank=", studentRank)

	// To delete all keys from map
	// Method 1
	for k, _ := range studentRank {
		delete(studentRank, k)
	}
	fmt.Println("After delete all , studentRank=", studentRank)

	// Method 2 - recreating map will reallocate memory by Garbage collector
	studentRank = make(map[string]int)
	fmt.Println("After re-creating , studentRank=", studentRank)

}

/*
Assignment:
	Get 10 times values from runtime, store them in slice
	Return number of times, each value is repeating

*/ 
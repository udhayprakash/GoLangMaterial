package main

import "fmt"

func main() {
	// Method 1 - Declaring, and  initializing the map
	studentRank1 := map[string]int{}
	fmt.Println("\nstudentRank1      = ", studentRank1)    // map[]
	fmt.Println("len(studentRank1) = ", len(studentRank1)) // 0

	// Method2 - Declaring, and  initializing the map
	studentRank := make(map[string]int)
	fmt.Println("\nstudentRank      = ", studentRank)    // map[]
	fmt.Println("len(studentRank) = ", len(studentRank)) // 0

	// To add/update key-value pairs
	studentRank["Suresh"] = 1
	studentRank["Mahesh"] = 3
	studentRank["Ganesh"] = 2

	fmt.Println("\nstudentRank      = ", studentRank)    // map[Ganesh:2 Mahesh:3 Suresh:1]
	fmt.Println("len(studentRank) = ", len(studentRank)) // 3

	// To overwrite all key-value pairs
	studentRank = map[string]int{"Apparao": 3, "subbarao": 1, "ramarao": 2}
	fmt.Println("\nstudentRank      = ", studentRank)    // map[Apparao:3 ramarao:2 subbarao:1]
	fmt.Println("len(studentRank) = ", len(studentRank)) // 3

	// maps are indexed by keys, not position
	fmt.Println("\nstudentRank[\"subbarao\"]=", studentRank["subbarao"]) // 1
	fmt.Println("studentRank[\"superman\"]=", studentRank["superman"])   // 0

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

	input - 11 44 11 66 11 44 33 78

	output - 11: 3, 44:2, 66: 1, 33:1, 78:1

	HINT: isKeyPresent
*/

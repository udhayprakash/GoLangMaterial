package main

import (
	"fmt"
	"reflect"
)

func main() {
	primesArray := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(reflect.TypeOf(primesArray),
		reflect.TypeOf(primesArray).Kind(),
		primesArray)
	// [6]int array [2 3 5 7 11 13]

	var slice1 []int = primesArray[1:4]
	fmt.Println(reflect.TypeOf(slice1),
		reflect.TypeOf(slice1).Kind(),
		slice1)
	// []int slice [3 5 7]

	// Indexing slices
	fmt.Println("slice1[0]            :", slice1[0])
	fmt.Println("slice1[len(slice1)-1]:", slice1[len(slice1)-1])

	alphabets := [...]string{1: "A", 2: "B", 3: "C", 4: "D"}
	fmt.Printf("\n\nalphabets\t\t\t\t : %v\n type\t\t\t\t : %T\n", alphabets, alphabets)

	fmt.Println("alphabets[0]            :", alphabets[0])             // ''
	fmt.Println("alphabets[1]            :", alphabets[1])             // A
	fmt.Println("alphabets[len(alphabets)-1]:", alphabets[len(alphabets)-1]) //D

	alphabets1 := []string{"A", "B", "C", "D"}
	fmt.Printf("\n\nalphabets1\t\t\t\t : %v\n type\t\t\t\t : %T\n", alphabets1, alphabets1)

	fmt.Println("alphabets1[0]            :", alphabets1[0])               // A
	fmt.Println("alphabets1[1]            :", alphabets1[1])               // B
	fmt.Println("alphabets1[len(alphabets1)-1]:", alphabets1[len(alphabets1)-1]) // D

}

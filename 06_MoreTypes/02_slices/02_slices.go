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

	alphas := [...]string{1: "A", 2: "B", 3: "C", 4: "D"}
	fmt.Printf("\n\nalphas\t\t\t\t : %v\n type\t\t\t\t : %T\n", alphas, alphas)

	fmt.Println("alphas[0]            :", alphas[0])             // ''
	fmt.Println("alphas[1]            :", alphas[1])             // A
	fmt.Println("alphas[len(alphas)-1]:", alphas[len(alphas)-1]) //D

	alphas1 := []string{"A", "B", "C", "D"}
	fmt.Printf("\n\nalphas1\t\t\t\t : %v\n type\t\t\t\t : %T\n", alphas1, alphas1)

	fmt.Println("alphas1[0]            :", alphas1[0])               // A
	fmt.Println("alphas1[1]            :", alphas1[1])               // B
	fmt.Println("alphas1[len(alphas1)-1]:", alphas1[len(alphas1)-1]) // D

}

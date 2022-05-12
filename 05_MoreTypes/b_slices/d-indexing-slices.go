package main

import (
	"fmt"
	"reflect"
)

func main() {
	primesArray := [6]int{2, 3, 5, 7, 11, 13}
	// indexing -         0  1  2  3  4   5
	fmt.Printf("Type-%T   value-%[1]v\n", primesArray)
	fmt.Println(reflect.TypeOf(primesArray).Kind()) // array

	var slice1 []int = primesArray[1:4] // [3, 5, 7]
	fmt.Printf("Type-%T   value-%[1]v\n", slice1)
	fmt.Println(reflect.TypeOf(slice1).Kind()) // slice

	// mutability check
	fmt.Println("slice1[1]=", slice1[1]) // 5
	slice1[1] = 555

	// Changes in slices are reflected back to array
	fmt.Println("slice1     =", slice1)      // [3 555 7]
	fmt.Println("primesArray=", primesArray) // [2 3 555 7 11 13]

	fmt.Println("slice1[1:3]=", slice1[1:3])
	fmt.Println()

	// slice1[1:3] = []int{55555, 77777}
	// cannot assign to slice1[1:3]

	// Try changing in array at that location, and check in slice
	fmt.Println("primesArray[1]=", primesArray[1]) // 3
	primesArray[1] = 333
	fmt.Println("slice1     =", slice1)      //  [333 555 7]
	fmt.Println("primesArray=", primesArray) //  [2 333 555 7 11 13]

	// Indexing slices
	fmt.Println("slice1[0]            :", slice1[0])             // 3
	fmt.Println("slice1[len(slice1)-1]:", slice1[len(slice1)-1]) // 7

	alphas := [...]string{1: "A", 2: "B", 4: "C", 5: "D"}
	fmt.Printf("\n\nalphas\t\t\t\t : %v\n type\t\t\t\t : %T\n", alphas, alphas)

	fmt.Println("alphas[0]            :", alphas[0])             // ''
	fmt.Println("alphas[1]            :", alphas[1])             // A
	fmt.Println("alphas[len(alphas)-1]:", alphas[len(alphas)-1]) //D

	alphas1 := []string{1: "A", 2: "B", 3: "C", 4: "D"}
	fmt.Printf("\n\nalphas1\t\t\t\t : %v\n type\t\t\t\t : %T\n", alphas1, alphas1)

	fmt.Println("alphas1[0]            :", alphas1[0])               // A
	fmt.Println("alphas1[1]            :", alphas1[1])               // B
	fmt.Println("alphas1[len(alphas1)-1]:", alphas1[len(alphas1)-1]) // D

}

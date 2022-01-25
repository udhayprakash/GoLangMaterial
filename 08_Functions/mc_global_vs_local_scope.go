package main

import "fmt"

func modifyArray(arr [10]int) { // cala by value
	arr[2] = 222
}

func modifySlice(mys []int) { // call by reference
	mys[2] = 222
}

func main() {
	var array1 [10]int
	fmt.Println("Array1 = ", array1) // [0 0 0 0 0 0 0 0 0 0]

	modifyArray(array1)
	fmt.Println("Array1 = ", array1) // [0 0 0 0 0 0 0 0 0 0]

	slice1 := []int{10, 20, 30, 40}
	fmt.Println("slice1 =", slice1) // [10 20 30 40]

	modifySlice(slice1)
	fmt.Println("slice1 =", slice1) // [10 20 222 40]

}

// Assignmengt: chedck for maps and structs too

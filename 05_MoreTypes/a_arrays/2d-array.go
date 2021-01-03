package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	var rows, columns int

	fmt.Println("Creating a 2 dimensional array of integer type ")
	fmt.Print("Enter number of rows : ")
	fmt.Scan(&rows)

	fmt.Print("Enter number of columns : ")
	fmt.Scan(&columns)

	// allocate memory for 2 dimension array of integer type with
	// [][]int
	// and use number of rows to determine the size
	array := make([][]int, rows)

	// followed by columns
	for i := range array {
		array[i] = make([]int, columns)
	}

	array[0][0] = 0 // in progamming world, starting counting from zero

	array[rows-1][columns-1] = rows * columns // assign value

	// print all assigned values by rows

	for r := range array {
		fmt.Printf("row(%d)  -  %d\n", r, array[r])
	}

	// trigger manual garbage collection to reclaim allocated memory
	array = nil
	debug.FreeOSMemory()
}

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

	// allocate memory for 2 dimension slice of integer type with
	// [][]int
	// and use number of rows to determine the size
	myslice := make([][]int, rows)

	for i := range myslice {
		myslice[i] = make([]int, columns)
	}
	fmt.Println("myslice = ", myslice)

	myslice[0][0] = rows * columns
	myslice[rows-1][columns-1] = rows * columns
	fmt.Println("myslice = ", myslice)

	for r := range myslice {
		fmt.Printf("row(%d)  -  %d\n", r, myslice[r])
	}

	// trigger manual garbage collection to reclaim allocated memory
	myslice = nil
	debug.FreeOSMemory()
}

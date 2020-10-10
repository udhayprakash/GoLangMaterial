package main

import "fmt"

/*
When array elements of an array are arrays,
then it’s called a multi-dimensional array.

*/

func main(){
	a1 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println("a1:", a1) // [[1 2] [3 4] [0 0]]
	fmt.Println()

	a2 := [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println("a2:", a2)
	fmt.Println()

	// Iterating multi-dimensional arrays
	for index1, childArray := range a2 {
		fmt.Println("At index ", index1, " we have", childArray)
		for _, element := range childArray {
			fmt.Printf("\telement: %d\n", element)
		}
	}
}
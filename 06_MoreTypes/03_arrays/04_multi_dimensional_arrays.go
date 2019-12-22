package main

import "fmt"

// When array elements of an array are arrays,
// then it’s called a multi-dimensional array.

func main() {
	a1 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println("a1:", a1)
	fmt.Println()

	a2 := [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println("a2:", a2)

	// Iterating multi-dimensional arrays
	for _, childArray := range a2 {
		for _, element := range childArray {
			fmt.Printf(element)
		}
	}
}


package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	fmt.Println("To get both index & values ...")
	for i, v := range pow {
		fmt.Printf("\t2 * %d = %3d\n", i, v)
	}

	fmt.Println("\nTo get only index ...")
	for index := range pow {
		fmt.Print(index, " ")
	}

	fmt.Println("\nTo get only values ...")
	for _, val := range pow {
		fmt.Print(val, " ")
	}

}

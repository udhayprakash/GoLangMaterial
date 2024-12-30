package main

import "fmt"

func length(args ...int) int {
	return len(args)
}

func main() {
	// slice
	vs := []int{24, 55, 88, 99, 4, 2}
	fmt.Println(length(vs...))


	vs = []int{24, 55,}
	fmt.Println(length(vs...))

	
	vs = []int{24}
	fmt.Println(length(vs...))

	vs = []int{}
	fmt.Println(length(vs...))
	fmt.Println()

	// direct values
	fmt.Println(length(12))
	fmt.Println(length(24, 55, 88, 99, 4, 2))

	// fmt.Println(length(213.213, 213.213. 213))

	// myarray := [4]int{12, 33, 45, 676}
	// fmt.Println(length(myarray...))
	// cannot use myarray (variable of type [4]int) as []int value in argument to length

}

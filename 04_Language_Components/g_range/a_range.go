package main

import "fmt"

func main(){
	// Looping over an Array
	nums := []int{11, 22, 33, 44, 55, 66, 77}
	for range nums {
		fmt.Println("- in - loop")
	}
	fmt.Println()

	for num := range nums {
		fmt.Println("Loop Index =", num)
	}
	fmt.Println()

	summation := 0
	for _, num := range nums {
		fmt.Println("num =", num)
		summation += num
	}
	fmt.Println("summation=", summation)


}

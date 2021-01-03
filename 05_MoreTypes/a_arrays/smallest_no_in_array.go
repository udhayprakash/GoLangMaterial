package main

import "fmt"

func main() {
	arr := []uint{
		28, 33, 16,
		7, 5, 88,
	}

	min := arr[0] // assume first value is the smallest

	for _, value := range arr {
		if value < min {
			min = value // found another smaller value, replace previous value in min
		}
	}

	fmt.Println("The smallest value is : ", min)
}

package main

import "fmt"

func main() {
	var numbers [10]int // array

	// // WRONG APPROACH
	// fmt.Println("Enter 10 numbers:")
	// _, err := fmt.Scanf("%v", &numbers)
	// if err != nil {
	// 	fmt.Println(err) // can't scan type: *[10]int
	// } else {
	// 	fmt.Println("numbers=", numbers)
	// }

	// correct way
	for index := 0; index < len(numbers); index++ {
		fmt.Printf("Enter %d position value:", index)
		fmt.Scanln(&numbers[index])
	}
	fmt.Println("numbers=", numbers)

}

// Assignment - try the same using range

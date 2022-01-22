package main

import "fmt"

func main() {
	var primeNumbers = []int{2, 5, 7}
	fmt.Println("Initially          		- ", primeNumbers)

	// appending an element
	primeNumbers = append(primeNumbers, 11)
	fmt.Println("after appending 11 		- ", primeNumbers)

	// appending Another Array of single Element
	primeNumbers = append(primeNumbers, []int{12}...)
	fmt.Println("after appending []int{12} 	- ", primeNumbers)

	// appending Another Array of multiple Elements
	primeNumbers = append(primeNumbers, []int{12, 13}...)
	fmt.Println("after appending []int{12, 13}	- ", primeNumbers)

	// appending Elements loosely
	primeNumbers = append(primeNumbers, 14, 15, 16)
	fmt.Println("after appending 14, 15, 16 loosely- ", primeNumbers)

	// Truncating the array
	primeNumbers = primeNumbers[4:9]
	fmt.Println("After truncating to primeNumbers[4:9] - ", primeNumbers)

}
 
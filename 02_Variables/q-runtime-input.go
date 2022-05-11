package main

import "fmt"

func main() {
	var userInput float32

	fmt.Printf("Please enter a floating point number: ")

	// _, err := fmt.Scanf("%f", &userInput) // %d
	_, err := fmt.Scan(&userInput) 
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("userInput = ", userInput)
	}

}
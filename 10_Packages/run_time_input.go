package main

import "fmt"

func main() {
	var userInput float32
	fmt.Printf("Please enter a floating point number: ")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Printf("An error occurred when trying to scan user input:\n\t%s", err)
	} else {
		fmt.Printf("The truncated integer value of your input is %d", int(userInput))
	}
}

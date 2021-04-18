package main

import "fmt"

func main() {

	var result int = 12
	if result == 12 {
		fmt.Println("result is 12")
	}

	if result != 12 {
		fmt.Println("result is NOT 12")
	}

	if !(result != 12) {
		fmt.Println("result is NOT 12")
	}

}

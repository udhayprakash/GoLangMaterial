package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input a string:")
	if scanner.Scan() {
		userInputLowerCase := strings.ToLower(scanner.Text())
		if strings.HasPrefix(userInputLowerCase, "i") ||
			strings.HasPrefix(userInputLowerCase, "a") ||
			strings.HasPrefix(userInputLowerCase, "n") {
			fmt.Println("found")
		} else {
			fmt.Println("NOT found")
		}

	} else {
		fmt.Printf("Error Occurred: %s", scanner.Err())
	}
}

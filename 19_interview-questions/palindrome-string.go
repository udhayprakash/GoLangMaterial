package main

import "fmt"

func isPalindrome(givenString string) bool {
	for i := 0; i < len(givenString)/2; i++ {
		if givenString[i] != givenString[len(givenString)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)

	result := isPalindrome(str)
	if result == true {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

package main

import (
	"fmt"
	"regexp"
	"unicode"
)

// Sample 01: “never a foot too far, even.” → true
// Sample 02: “,,,,never a foot too far, even.” → true
// Sample 03: “”hello world”” → false
// Sample 04: “”racecar”” → true

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

	result = isPalindrome2(str)
	if result == true {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}

	result = isPalindrome3(str)
	if result == true {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

func isPalindrome2(sentence string) bool {
	/*
		Time and Space Complexity
			Time Complexity: O(N * ((N-1) + (N-1))). Given N is number of characters in the string.
			Space Complexity: O(1). No auxiliary memory space used.
	*/
	left := 0
	right := len(sentence) - 1

	for {
		if left > right {
			break
		}
		for {
			if unicode.IsLetter(rune(sentence[left])) == true {
				break
			}
			left += 1
		}
		for {
			if unicode.IsLetter(rune(sentence[right])) == true {
				break
			}
			right -= 1
		}

		if sentence[left] != sentence[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func isPalindrome3(sentence string) bool {
	/*
		Time and Space Complexity
			Time Complexity: O(N+N). Given N is number of characters in the string.
			Space Complexity: O(1). No auxiliary memory space used.
	*/
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	processedSentence := reg.ReplaceAllString(sentence, "")
	left := 0
	right := len(processedSentence) - 1
	for {
		if left > right {
			break
		}
		if processedSentence[left] != processedSentence[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

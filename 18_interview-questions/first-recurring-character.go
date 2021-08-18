package main

import "fmt"

/*
Problem: Given a string, return the first recurring character in it,
 or null if there is no recurring character.

For example, given the string "acbbac", return "b".
			 Given the string "abcdef", return null.
*/
func main() {
	fmt.Println("acbbac", getFirstRecurringChar("acbbac"))
	fmt.Println("abcdef", getFirstRecurringChar("abcdef"))
}

func getFirstRecurringChar(givenString string) string {
	chars := make(map[string]int)
	for _, char := range givenString {
		_, isExits := chars[string(char)]
		if isExits == true {
			return string(char)
		}
		chars[string(char)] = 1
	}
	return ""
}

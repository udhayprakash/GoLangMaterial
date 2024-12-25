/*
Reverse vowels in a given string
Given a string, your task is to reverse only the vowels of the string.

Examples:
  Input : programming
  Output : prigrammong
  
  Input : hello world
  Output : hollo werld
  
  Input : production grade golang code
  Output : predoctaon greda golong cudo
*/
package main

import (
	"fmt"
	"strings"
)

func reverseStringVowels(givenstr string) string {
	vowels := []rune{}
	indices := []int{}
	for i, r := range givenstr {   // O(N)
		if strings.ContainsRune("aeiouAEIOU", r) {
			vowels = append(vowels, r)
			indices = append(indices, i)
		}
	}
	result := []rune(givenstr)
	for _, i := range indices {
		result[i] = vowels[len(vowels)-1]
		vowels = vowels[:len(vowels)-1]
	}
	return string(result)
}

func main() {
	fmt.Println(reverseStringVowels("programming"))
	fmt.Println(reverseStringVowels(" hello world"))

}

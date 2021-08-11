package main

import (
	"fmt"
	"strings"
)

/*
Given a string of words delimited by spaces, reverse the words in string.
For example, given "hello world here", return "here world hello"

Follow-up: given a mutable string representation, can you perform this operation in-place?
*/
func main() {
	fmt.Println("hello world here:", wordReverse("hello world here"))        // here world hello
	fmt.Println("hello world here:", wordReverseInPlace("hello world here")) // here world hello
	fmt.Println("hello world here:", wordReverse2("hello world here"))       // olleh dlrow ereh
}

func wordReverse(givenSentence string) string {
	reverseSentence := ""
	for _, word := range strings.Fields(givenSentence) {
		reverseSentence = word + " " + reverseSentence
	}
	return reverseSentence
}

// TODO
func wordReverseInPlace(givenSentence string) string {
	reverseSentence := ""
	for index, word := range strings.Fields(givenSentence) {
		fmt.Println(index, word)
	}
	return reverseSentence
}

func wordReverse2(givenSentence string) string {
	reverseSentence := ""
	reverseWord := ""
	for _, word := range strings.Fields(givenSentence) {
		for i := len(word) - 1; i >= 0; i-- {
			reverseWord += string(word[i])
		}
		// fmt.Println(word, reverseWord)
		if reverseSentence == "" {
			reverseSentence = reverseWord
		} else {
			reverseSentence += " " + reverseWord
		}
		reverseWord = ""
	}
	return reverseSentence
}

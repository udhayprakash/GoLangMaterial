package main

import (
	"fmt"
	"sort"
	"strings"
)

func reverse(s string) string {

	words := strings.Fields(s) // tokenize each words from input string
	totalLength := len(words)

	// reverse the order(no sorting!)
	for i, j := 0, totalLength-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// return the reversed words
	return strings.Join(words, " ")
}

func main() {
	reversed := reverse("apple durian kiwi banana")

	fmt.Println(reversed)

	reverseString := sort.Reverse("apple durian kiwi banana")
	fmt.Println(reverseString)
}

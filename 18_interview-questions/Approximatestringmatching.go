package main

import (
	"fmt"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

func main() {

	// find
	input := []string{"example", "help", "assistance", "existence"}

	fuzzyMatches := fuzzy.Find("ex", input)

	fmt.Println("Matches found : ", fuzzyMatches)
}

// Ref: https://en.wikipedia.org/wiki/Approximatestringmatching

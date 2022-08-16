package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("example - existence	-> %0.5f\n", JaroWinklerDistance("example", "existence"))
	fmt.Printf("feel 	- fill 		-> %0.5f\n", JaroWinklerDistance("feel", "fill"))
	fmt.Printf("octopus -         	-> %0.5f\n", JaroWinklerDistance("octopus", ""))
	fmt.Printf("stick 	- stix 		-> %0.5f\n", JaroWinklerDistance("stick", "stix"))
	fmt.Printf("top 	- stop 		-> %0.5f\n", JaroWinklerDistance("top", "stop"))
	fmt.Printf("tick 	- lick 		-> %0.5f\n", JaroWinklerDistance("tick", "lick"))
	fmt.Printf("golang 	- Golang 	-> %0.5f\n", JaroWinklerDistance("golang", "golang"))
}

// JaroWinklerDistance - calculate and return the Jaro-Winkler distance
func JaroWinklerDistance(s1, s2 string) float64 {

	s1Matches := make([]bool, len(s1)) // |s1|
	s2Matches := make([]bool, len(s2)) // |s2|

	var matchingCharacters = 0.0
	var transpositions = 0.0

	// sanity checks

	// return 0 if either one is empty string
	if len(s1) == 0 || len(s2) == 0 {
		return 0 // no similarity
	}

	// return 1 if both strings are empty
	if len(s1) == 0 && len(s2) == 0 {
		return 1 // exact match
	}

	if strings.EqualFold(s1, s2) { // case insensitive
		return 1 // exact match
	}

	// Two characters from s1 and s2 respectively,
	// are considered matching only if they are the same and not farther than
	// [ max(|s1|,|s2|) / 2 ] - 1
	matchDistance := len(s1)
	if len(s2) > matchDistance {
		matchDistance = len(s2)
	}
	matchDistance = matchDistance/2 - 1

	// Each character of s1 is compared with all its matching characters in s2
	for i := range s1 {
		low := i - matchDistance
		if low < 0 {
			low = 0
		}
		high := i + matchDistance + 1
		if high > len(s2) {
			high = len(s2)
		}
		for j := low; j < high; j++ {
			if s2Matches[j] {
				continue
			}
			if s1[i] != s2[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matchingCharacters++
			break
		}
	}

	if matchingCharacters == 0 {
		return 0 // no similarity, exit early
	}

	// Count the transpositions.
	// The number of matching (but different sequence order) characters divided by 2 defines the number of transpositions
	k := 0
	for i := range s1 {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			transpositions++ // increase transpositions
		}
		k++
	}
	transpositions /= 2

	weight := (matchingCharacters/float64(len(s1)) + matchingCharacters/float64(len(s2)) + (matchingCharacters-transpositions)/matchingCharacters) / 3

	//  the length of common prefix at the start of the string up to a maximum of four characters
	l := 0

	// is a constant scaling factor for how much the score is adjusted upwards for having common prefixes.
	//The standard value for this constant in Winkler's work is {\displaystyle p=0.1}p=0.1
	p := 0.1

	// make it easier for s1[l] == s2[l] comparison
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if weight > 0.7 {
		for (l < 4) && s1[l] == s2[l] {
			l++
		}

		weight = weight + float64(l)*p*(1-weight)
	}

	return weight
}

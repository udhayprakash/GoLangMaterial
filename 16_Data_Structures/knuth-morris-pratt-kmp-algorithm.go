// Knuth–Morris–Pratt (KMP) in Golang
package main

import (
	"fmt"
)

// The Knuth–Morris–Pratt string searching algorithm (or KMP algorithm) searches for occurrences of a "substring" within a main "string" by employing the observation that when a mismatch occurs, the word itself embodies sufficient information to determine where the next match could begin, thus bypassing re-examination of previously matched characters.
const (
	PatternSize int = 100
)

func SearchNext(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[len(retSlice)-1]
	}

	return -1
}

func SearchString(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[0]
	}

	return -1
}

func KMP(haystack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var mpNext [PatternSize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}

func preKMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [PatternSize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}

func main() {
	fmt.Println("Search First Position String:\n")
	fmt.Println(SearchString("cocacola", "co"))
	fmt.Println(SearchString("Australia", "lia"))
	fmt.Println(SearchString("cocacola", "cx"))
	fmt.Println(SearchString("AABAACAADAABAABA", "AABA"))

	fmt.Println("\nSearch Last Position String:\n")
	fmt.Println(SearchNext("cocacola", "co"))
	fmt.Println(SearchNext("Australia", "lia"))
	fmt.Println(SearchNext("cocacola", "cx"))
	fmt.Println(SearchNext("AABAACAADAABAABAAABAACAADAABAABA", "AABA"))
}

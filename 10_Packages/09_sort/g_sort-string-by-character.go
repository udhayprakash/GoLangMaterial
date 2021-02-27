package main

import (
	"fmt"
	"sort"
)

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	fmt.Println(sortStringByCharacter("listen"))
	fmt.Println(sortStringByCharacter("silent"))

	if sortStringByCharacter("listen") == sortStringByCharacter("silent") {
		fmt.Println("Both are anagrams")
	}
}

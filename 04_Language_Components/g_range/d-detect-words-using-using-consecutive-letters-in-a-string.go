package main

import (
	"fmt"
	"strings"
)

func main() {

	textString := `Bookkeeper and bookkeeping and sweet-toothed
	have three consecutive sets of double letters.Others are 
	deer-reeve, feed-door, heel-loop, hoof-footed, hoot-toot, 
	keek-keek, Soonnee, toot-toot, veneer-room, and wood-deer.`
	fmt.Println(textString)

	words := strings.Split(textString, " ")
	fmt.Println(words)

	// there could be more, but let's start with aa ee oo first.
	doubleLetters := []string{"aa", "ee", "oo"}
	fmt.Println(doubleLetters)

	for i := 0; i < len(doubleLetters); i++ {
		for _, word := range words {
			if strings.Contains(word, doubleLetters[i]) {
				fmt.Println(word)
			}
		}
	}

}

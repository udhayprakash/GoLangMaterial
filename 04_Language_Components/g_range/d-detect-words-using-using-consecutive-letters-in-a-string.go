package main

import (
	"fmt"
	"strings"
)

func main() {

	textString := "Bookkeeper and bookkeeping and sweet-toothed have three consecutive sets of double letters.Others are deer-reeve, feed-door, heel-loop, hoof-footed, hoot-toot, keek-keek, Soonnee, toot-toot, veneer-room, and wood-deer."

	words := strings.Split(textString, " ")

	// there could be more, but let's start with aa ee oo first.
	doubleLetters := []string{"aa", "ee", "oo"}

	fmt.Println(textString)

	for i := 0; i < len(doubleLetters); i++ {

		for _, word := range words {
			if strings.Contains(word, doubleLetters[i]) {
				fmt.Println("[" + word + "] using consecutive letters")
			}
		}
	}

}

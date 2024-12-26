package main

import "fmt"

func main() {
	// words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	words := []string{"", "1", "22", "33", "4444", "55555", "666666", "7777777", "999999999"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9: // skipped
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

// NOTE: default case is optional
// In switch case,   -- break isn't needed like in other languages
//  -- if no expression provided, go checks for the first case that evals to true

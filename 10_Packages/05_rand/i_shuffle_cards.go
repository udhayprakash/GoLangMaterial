package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suites = [4]string{"D", "S", "C", "H"}
var ranks = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func createDeck() []string {
	deck := []string{}
	for _, s := range suites {
		for _, r := range ranks {
			deck = append(deck, s+r)
		}
	}
	return deck
}

func shuffleDeck(a []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func main() {
	deck := createDeck()
	fmt.Printf("Initial deck :\n %+v\n", deck)

	deck = shuffleDeck(deck)
	fmt.Printf("Shuffled deck:\n %+v\n", deck)

	hands := deck[:5]
	fmt.Printf("Hands        :\n %+v\n", hands)

}

package main

import (
	"fmt"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

func main() {

	input := []string{"example", "help", "assistance", "existence"}

	rankMatches := fuzzy.RankFind("ex", input)

	for _, rank := range rankMatches {
		fmt.Println("Source : ", rank.Source, " Word :", rank.Target, " Distance : ", rank.Distance)
	}

}

/*
Ref:
https://en.wikipedia.org/wiki/Levenshtein_distance

https://github.com/renstrom/fuzzysearch
*/

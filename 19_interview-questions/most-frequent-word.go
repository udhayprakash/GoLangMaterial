package main

import (
	"fmt"
	"regexp"
)

func main() {
	sentence1 := "I felt happy because I saw the others were happy and because I knew I should feel happy, but I wasn't really happy"
	fmt.Println(`GetMostFrequentWord(sentence1, []string{"i"})=`, GetMostFrequentWord(sentence1, []string{"i"}))

	sentence1 = "I felt Happy because i, saw the others were happy. And because i knew i should feel Happy, but I wasn't really happy!"
	fmt.Println(`GetMostFrequentWord(sentence1, []string{"i"})=`, GetMostFrequentWord(sentence1, []string{"i"}))
	fmt.Println(`GetMostFrequentWord(sentence1, []string{"i", "I"})=`, GetMostFrequentWord(sentence1, []string{"i", "I"}))
	fmt.Println(`GetMostFrequentWord(sentence1, []string{"i", "I", "Happy"})=`, GetMostFrequentWord(sentence1, []string{"i", "I", "Happy"}))
}

func GetMostFrequentWord(sent string, excludeList []string) string {
	exclusionMap := make(map[string]int)
	for _, str := range excludeList {
		exclusionMap[str] = 1
	}
	// words := strings.Split(sent, " ")
	words := regexp.MustCompile("[^a-zA-Z0-9]+").Split(sent, -1)

	wrdFreq := make(map[string]int)
	for _, word := range words {
		_, isPresent := exclusionMap[word]
		if isPresent == false {
			wrdFreq[word] += 1
		}
	}
	highCount, highWord := 0, ""
	for wrd, count := range wrdFreq {
		if count > highCount {
			highCount, highWord = count, wrd
		}
	}
	fmt.Println("\n", wrdFreq)
	return highWord
}

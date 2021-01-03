package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func main() {
	originalText := `Psssst! Hey you! Yes! You there! This is my secret message to you`

	fmt.Println("Original: ", originalText)

	// now we want to randomly mark some words to be replaced
	var scrambledText, unscrambledText []string

	// convert text to slice
	textSlice := strings.Split(originalText, " ")
	//fmt.Println(textSlice)

	// prepare the dice
	dice := []int{1, 2, 3, 4, 5, 6}
	rand.Seed(time.Now().UnixNano())

	// prepare the translation map
	translateBackMap := map[string]string{}
	var randomStringLength = 8

	for _, v := range textSlice {

		// roll our dice and if less than 3(or any number that you want)
		// mark the word. This is to ensure the words will be marked randomly
		rollTheDice := dice[rand.Intn(len(dice))]
		//fmt.Println("Rolled: ", rollTheDice)

		if rollTheDice <= 3 {
			// replace the words and insert into a translation map
			// so that we can translate back the cryptic replacement words back later
			replacement := randStr(randomStringLength, "alphanum")
			translateBackMap[replacement] = v
			scrambledText = append(scrambledText, replacement)
		} else {
			scrambledText = append(scrambledText, v)
		}
	}

	result := strings.Join(scrambledText, " ")

	fmt.Println("Scrambled: ", result)

	// now we want to translate back those cryptic replacement words to something readable
	// the idea here is to look for words with the size of 8 - the variable randomStringLength
	// if the length matches randomStringLength, then look for the replacement words in the
	// translateBackMap

	for _, v := range scrambledText {
		if len(v) == 8 {
			unscrambledText = append(unscrambledText, translateBackMap[v])
		} else {
			unscrambledText = append(unscrambledText, v)
		}
	}

	final := strings.Join(unscrambledText, " ")
	fmt.Println("Unscrambled: ", final)
}

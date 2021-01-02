package main

import (
	"fmt"
	"math/rand"
	"time"
)

//var vowels = []rune{'a', 'e', 'i', 'o', 'u'}
//var consonants = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n','p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}

func humanReadablePassword(alphabetSize, numberSize int) string {

	vowels := "aeiou"
	consonants := "bcdfghjklmnpqrstvwxyz"
	digits := "0123456789"

	prefixSize := alphabetSize / 2
	if alphabetSize%2 != 0 {
		prefixSize = int(alphabetSize/2) + 1
	}
	suffixSize := alphabetSize - prefixSize

	var prefixPart = make([]byte, prefixSize)

	for i := 0; i <= prefixSize-1; i++ {
		if i%2 == 0 {
			// use consonants
			prefixPart[i] = consonants[rand.Intn(len(consonants)-1)]
		} else {
			// use vowels
			prefixPart[i] = vowels[rand.Intn(len(vowels)-1)]
		}
	}

	var midPart = make([]byte, numberSize)

	// use digits
	for k, _ := range midPart {
		midPart[k] = digits[rand.Intn(len(digits))]
	}

	var suffixPart = make([]byte, suffixSize)

	for i := 0; i <= suffixSize-1; i++ {
		if i%2 == 0 {
			// use consonants
			suffixPart[i] = consonants[rand.Intn(len(consonants)-1)]
		} else {
			// use vowels
			suffixPart[i] = vowels[rand.Intn(len(vowels)-1)]
		}
	}

	return string(prefixPart) + string(midPart) + string(suffixPart)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("6 alphabets with 2 digits : ", humanReadablePassword(6, 2)) // best option
	fmt.Println("3 alphabets with 8 digits : ", humanReadablePassword(3, 8))
	fmt.Println("9 alphabets with 9 digits : ", humanReadablePassword(9, 9))

}

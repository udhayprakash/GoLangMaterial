package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Purpose: Password Generation
	- Random ASCII string with at least 1 digit and
        1 special character
*/
func main() {
	password := passwordGenerator(10)
	fmt.Println("password :", password)
}

func passwordGenerator(wordLength int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + digits +
		"abcdefghijklmnopqrstuvwxyz" + specials
	buf := make([]byte, wordLength)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < wordLength; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	for i := len(buf) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		buf[i], buf[j] = buf[j], buf[i]
	}
	str := string(buf)
	return str
}

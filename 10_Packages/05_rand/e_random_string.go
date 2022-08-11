package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genRandomUpperCaseString(strLen int) string {
	var result string
	for i := 0; i < strLen; i++ {
		result += string('A' + rune(rand.Intn('Z'-'A'+1)))
	}
	return result
}

func genRandomString(strLength int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	buf := make([]rune, strLength)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	return string(buf)
}

func main() {
	fmt.Println("genRandomUpperCaseString(10):", genRandomUpperCaseString(10)) // XVLBZGBAIC
	fmt.Println()

	// rand.Seed(time.Now().UnixNano()) -- adding it makes unpredictable
	fmt.Println("genRandomString(10):", genRandomString(10)) // BsD9ufNhpt
	fmt.Println("genRandomString(10):", genRandomString(10)) // fPnZUSLy7r
	rand.Seed(32)
	fmt.Println("\ngenRandomString(10):", genRandomString(10)) // jLFmRZphVf
	fmt.Println("genRandomString(10):", genRandomString(10))   // p1kar48Air
	fmt.Println()

	// case when rand.Seed in function
	fmt.Println("genRandomString2(10):", genRandomString2(10)) // unpredictable
	fmt.Println("genRandomString2(10):", genRandomString2(10)) // unpredictable
}

func genRandomString2(strLength int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials

	buf := make([]byte, strLength)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < strLength; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

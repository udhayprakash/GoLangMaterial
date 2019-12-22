package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rdString := genRandomString(10)
	fmt.Println("str :", rdString)

	fmt.Println("genRandomString(10):", genRandomString(10))
	rand.Seed(2123)
	fmt.Println("genRandomString(10):", genRandomString(10))
}

func genRandomString(strLength int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	buf := make([]rune, strLength)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	str := string(buf)
	return str
}

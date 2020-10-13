package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genRandomString(strLength int) string {

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	buf := make([]rune, strLength)
	for i := range buf{
		buf[i] = chars[rand.Intn(len(chars))]
		//fmt.Println("buf[i]:", buf[i])
	}
	//fmt.Println("buf        :", buf)
	//fmt.Println("string(buf):", string(buf))
	return string(buf)
}

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("genRandomString(10):", genRandomString(10))
	fmt.Println("genRandomString(10):", genRandomString(10))
	rand.Seed(32)
	fmt.Println("\ngenRandomString(10):", genRandomString(10))
	fmt.Println("genRandomString(10):", genRandomString(10))
}
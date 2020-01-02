package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Without seed, output will be same on re-execution
	rand.Seed(time.Now().UnixNano())
	c1 := 'a' + rand.Intn(26)
	fmt.Println("c1 =", c1)

	chars := []rune("ab⌘")
	c2 := chars[rand.Intn(len(chars))]
	fmt.Println("c2 =", c2)
}

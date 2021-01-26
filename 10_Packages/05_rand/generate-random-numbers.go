package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1000, 2000)
	fmt.Printf("Random Num: %d\n", randomNum)
}

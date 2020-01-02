package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	lowerBound := 10
	upperBound := 100
	rand.Seed(time.Now().UnixNano())
	num := lowerBound + rand.Intn(upperBound-lowerBound+1)

	fmt.Println("num =", num)
}

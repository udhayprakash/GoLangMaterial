package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Mersene Twister algorithms - PRNG algorithms
func main() {
	// To produce varying sequences, give it a seed that changes.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println("r1.Intn(100):", r1.Intn(100)) // -- changing
	fmt.Println("r1.Intn(100):", r1.Intn(100)) // -- changing

	s1 = rand.NewSource(12313)
	r1 = rand.New(s1)

	fmt.Println("\nr1.Intn(100):", r1.Intn(100)) // 49
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 74
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 1

	s1 = rand.NewSource(234)
	r1 = rand.New(s1)

	fmt.Println("\nr1.Intn(100):", r1.Intn(100)) // 8
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 89
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 43
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 44
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 46
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 49
	fmt.Println("r1.Intn(100):", r1.Intn(100))   // 34
	fmt.Println()

	fmt.Println("r1.Float32():", r1.Float32())   // 0.93458873
	fmt.Println("r1.Float32():", r1.Float32())   // 0.55352
	fmt.Println()

	fmt.Println("r1.Float64():", r1.Float64())   // 0.1776533225905847
	fmt.Println("r1.Float64():", r1.Float64())   // 0.7141261020990882
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// random integr between 0 & 100
	fmt.Println("rand.Intn(100):", rand.Intn(100))
	fmt.Println("rand.Intn(100):", rand.Intn(100))
	fmt.Println("rand.Intn(100):", rand.Intn(100))
	// between 0 to 25
	fmt.Println("rand.Intn(25) :", rand.Intn(25))

	// `rand.Float64` returns a `float64` `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println("\nrand.Float64():", rand.Float64())
	fmt.Println("rand.Float64():", rand.Float64())
	fmt.Println("rand.Float64():", rand.Float64())
	fmt.Println("rand.Float64():", rand.Float64())

	// This can be used to generate random floats in
	// other ranges, for example `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

}

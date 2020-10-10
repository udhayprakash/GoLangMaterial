package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("e^%d = %8.3f\n", x, math.Exp(float64(x)))
	}
}

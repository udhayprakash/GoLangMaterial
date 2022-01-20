package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const Avogadro = 6.02214129e23
	fmt.Println(Avogadro, reflect.TypeOf(Avogadro))

	var Avogadro2 = 6.02214129 * math.Pow(10, 23)
	fmt.Println(Avogadro2, reflect.TypeOf(Avogadro2))

	fmt.Println()

	positiveInfinity := math.Inf(0)
	fmt.Println(positiveInfinity, reflect.TypeOf(positiveInfinity)) // +Inf float64

	negativeInfinity := math.Inf(-98)
	fmt.Println(negativeInfinity, reflect.TypeOf(negativeInfinity)) // -Inf float64

	garbageValue := math.Inf
	fmt.Println(garbageValue, reflect.TypeOf(garbageValue)) // 0xc6c5a0 func(int) float64
	fmt.Println()

	x := 1
	fmt.Println("math.Exp(float64(x)):", math.Exp(float64(x)))
	x = 2
	fmt.Println("math.Exp(float64(x)):", math.Exp(float64(x)))
	x = 3
	fmt.Println("math.Exp(float64(x)):", math.Exp(float64(x)))

	for x := 0; x < 10; x++ {
		fmt.Printf("x = %d eA = %10.2f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
	// NaN - not-a-number value
}

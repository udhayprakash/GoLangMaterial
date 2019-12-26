package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	fmt.Println(Avogadro, reflect.TypeOf(Avogadro))
	fmt.Println(Planck, reflect.TypeOf(Planck))

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
	// NaN - not-a-number value


}

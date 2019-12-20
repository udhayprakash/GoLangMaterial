package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
)

func main() {
	var cNum complex64 = 3 + 4i
	fmt.Println("cNum                   =", cNum)
	fmt.Println("reflect.TypeOf(cNum)   =", reflect.TypeOf(cNum))
	fmt.Println("cmplx.Abs(3+4i)        =", cmplx.Abs(3+4i))
	fmt.Println("cmplx.Conj(3+4i)       =", cmplx.Conj(3+4i))

	fmt.Println("cmplx.Exp(1i*math.Pi)+1=", cmplx.Exp(1i*math.Pi)+1)


	r, theta := cmplx.Polar(2i)
	fmt.Printf("r: %.1f, θ: %.1f*π", r, theta/math.Pi)

}

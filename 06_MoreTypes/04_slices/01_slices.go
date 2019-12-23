package main

import (
	"fmt"
	"reflect"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, reflect.TypeOf(primes))

	var s []int = primes[1:4]
	fmt.Println(s, reflect.TypeOf(s))
}

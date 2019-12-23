package main

import (
	"fmt"
	"reflect"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, reflect.TypeOf(primes), reflect.TypeOf(primes).Kind())
	// [2 3 5 7 11 13] [6]int array

	var s []int = primes[1:4]
	fmt.Println(s, reflect.TypeOf(s), reflect.TypeOf(s).Kind())
	// [3 5 7] []int slice
}

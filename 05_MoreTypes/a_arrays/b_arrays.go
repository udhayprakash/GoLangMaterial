package main

import (
	"fmt"
	"reflect"
)

func main(){

	// comparision on arrays
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{2, 3, 1}
	a3 := [3]int{3, 1, 2}
	a4 := [...]int{1, 2, 3}

	fmt.Println("a1 == a2 :", a1 == a2) // false
	fmt.Println("a1 == a3 :", a1 == a3) // false
	fmt.Println("a1 == a4 :", a1 == a4) // true


	fmt.Printf("a1 is of type : %T\n", a1)              // [3]int
	fmt.Println("a1 is of type :", reflect.TypeOf(a1)) // [3]int
	fmt.Println()

	q := [3]int{1, 2, 3}
	fmt.Println("q =", q) // [1 2 3]
	q = [3]int{44, 55, 66}
	fmt.Println("q =", q) // [44 55 66]

	//q = [5]int{11, 22, 33, 44, 55} // cannot use [5]int literal (type [5]int) as type [3]int in assignment

	//q = [3]float64{4.4, 5.5, 6.6} // cannot use [3]float64 literal (type [3]float64) as type [3]int in assignment

}
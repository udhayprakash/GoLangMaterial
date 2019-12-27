package main

import (
	"fmt"
	"reflect"
)

func main() {
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

	// array size cant be changed
	//q := [3]int{1, 2, 3}
	//q=[4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

}

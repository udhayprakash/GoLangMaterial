package main

import "fmt"

func main() {
	var as []string
	fmt.Println("Initially,      as = ", as) // []

	// works with slice only
	// append(as, "Udhay") // Error: evaluated but not used
	as = append(as, "Udhay")
	fmt.Println("After append(), as = ", as) // [Udhay]

	var aNum []int
	fmt.Println("\nIntitially    , aNum = ", aNum) // []
	aNum = append(aNum, 123)
	fmt.Println("After append(), aNum = ", aNum) // [123]

	var aInter []interface{}

	fmt.Println("\nInitially, aInter=", aInter) // []
	aInter = append(aInter, 123)
	aInter = append(aInter, 123.5)
	aInter = append(aInter, "123.5")
	aInter = append(aInter, "u")
	aInter = append(aInter, `u`)
	aInter = append(aInter, "udhay")
	fmt.Println("After, aInter=", aInter) // [123 123.5 123.5 u u udhay]

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x) // [1 2 3 4 5 6]

	p := []int{1, 2, 3}
	q := []int{4, 5, 6}
	p = append(p, q...) // concatenation two arrays
	fmt.Println(p)      // [1 2 3 4 5 6]
	// Without that ..., it wouldn't compile because the
	// types would be wrong; q is not of type int.

}

package main

import "fmt"

func main() {
	usd := 76
	cand := 50

	fmt.Println("usd < cand  =", usd < cand)
	fmt.Println("usd <= cand =", usd <= cand)
	fmt.Println("usd > cand  =", usd > cand)
	fmt.Println("usd >= cand =", usd >= cand)
	fmt.Println("usd == cand =", usd == cand)

	fmt.Println("usd != cand =", usd != cand)
	// fmt.Println("usd <> cand =", usd <> cand)
	// syntax error: unexpected >, expecting expression

	fmt.Println(12 < 34) // true

	// fmt.Println(12 < 34 < 45)
	// invalid operation: 12 < 34 < 45 (mismatched types untyped bool and untyped int)

	fmt.Println(12 < 34)       // true
	fmt.Println(1.2 < 3.4)     // true
	fmt.Println("1.2" < "3.4") // true
	fmt.Println('1' < '3')     // true
	fmt.Println('1' < 3)       // false

	// string comparision
	fmt.Println("apple" < "apparao") // false

	fmt.Println(4 == 4.0) // true
	// var (
	// 	num3 int     = 4
	// 	num4 float32 = 4.0
	// )
	// fmt.Println("num3 == num4:", num3 == num4)
	// invalid operation: num3 == num4 (mismatched types int and float32)
}

// assignment - try for all other builtin data types

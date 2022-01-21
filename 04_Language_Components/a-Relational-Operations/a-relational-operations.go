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

	// string comparision
	fmt.Println("apple" < "apparao") // false

}

// assignment - try for all other builtin data types

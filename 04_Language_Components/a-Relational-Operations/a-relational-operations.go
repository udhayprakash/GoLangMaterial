package main

import "fmt"

func main() {
	usd := 71
	cand := 49

	fmt.Println("usd < cand  =", usd < cand)
	fmt.Println("usd <= cand =", usd <= cand)
	fmt.Println("usd > cand  =", usd > cand)
	fmt.Println("usd >= cand =", usd >= cand)
	fmt.Println("usd == cand =", usd == cand)
	fmt.Println("usd != cand =", usd != cand)

	// fmt.Println("usd <> cand =", usd <> cand)
	// syntax error: unexpected >, expecting expression
	
	fmt.Println(12 < 34) // true
	// fmt.Println(12 < 34 < 45) // InValid Syntax 
}

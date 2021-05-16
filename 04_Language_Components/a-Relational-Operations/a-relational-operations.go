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
}

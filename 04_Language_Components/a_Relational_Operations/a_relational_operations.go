package main

import "fmt"

func main() {
	usd := 74
	cand := 48

	fmt.Println("usd < cand  =", usd < cand)
	fmt.Println("usd <= cand =", usd <= cand)
	fmt.Println("usd > cand  =", usd > cand)
	fmt.Println("usd >= cand =", usd >= cand)
	fmt.Println("usd == cand =", usd == cand)
	fmt.Println("usd != cand =", usd != cand)
	//fmt.Println("usd <> cand =", usd <> cand) syntax error: unexpected >, expecting expression

}

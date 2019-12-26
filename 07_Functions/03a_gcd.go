package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y // tuple assignment
	}
	return x
}

func main() {
	fmt.Println("gcd(4, 12)  :", gcd(4, 12))
	fmt.Println("gcd(22, 50) :", gcd(22, 50))
	fmt.Println("gcd(-20, 50):", gcd(-20, 50))
}

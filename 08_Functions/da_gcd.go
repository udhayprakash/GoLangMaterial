package main

import "fmt"

// Function Definitions
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y // tuple assignment
	}
	return x
}

func main() {
	res := gcd(4, 12)
	fmt.Println("res =", res)

	fmt.Println("gcd(4, 12)  :", gcd(4, 12))
	fmt.Println("gcd(22, 50) :", gcd(22, 50))
	fmt.Println("gcd(-20, 50):", gcd(-20, 50))
}

// Assignment: Enhance this solution by passing values from a []int{}
//             and calling in for loop

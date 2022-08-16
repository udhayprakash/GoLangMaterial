package main

import "fmt"

/*
Problem: Write a function that takes in a number of
ounces and returns the smallest number of bars of
soap required to fulfill the order.
There are 3 sizes of soap bars: 2 oz, 5 oz, 10 oz.

Ex: 48 ounces -> (4 * 10 oz) + (1 * 5 oz) + (1 * 2 oz) | (1 left)
*/
func main() {
	GetMinSoups(1)
	GetMinSoups(2)
	GetMinSoups(5)
	GetMinSoups(7)
	GetMinSoups(10)
	GetMinSoups(48)
}

func GetMinSoups(ounces int) {
	soups := 0
	values := [3]int{10, 5, 2}
	fmt.Println("\nounces:", ounces)

	for _, val := range values {
		fmt.Println("\t", ounces/val, "*", val, "oz")
		soups += ounces / val
		ounces -= (ounces / val) * val
		if ounces == 0 {
			break
		}
	}
	fmt.Println("Total soups count:", soups)
}

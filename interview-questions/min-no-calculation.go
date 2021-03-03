package main

import "fmt"

func main() {
	r := GetMinSoups(48)
	fmt.Println(r)
}

func GetMinSoups(ounces int) int {
	soups := 0
	values := [3]int{10, 5, 2}
	for _, val := range values {
		fmt.Println(ounces, val, ounces/val)

		if ounces%val == 0 {

		}
	}
	return soups
}

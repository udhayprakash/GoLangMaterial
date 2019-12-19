package main

import "fmt"

func add(x int, y int ) int {
	var output = x + y
	return output
}

func add1(x int, y int ) (output int) {
	output = x + y
	return output
}

func divMod(x, y int) (q int, r int) {
	q = x /y
	r = x % y
	return q, r
}

func main() {
	num1 := 100
	num2 := 200

	result := add(num1, num2)
	fmt.Println("result   =", result)

	result1 := add1(num1, num2)
	fmt.Println("result1  =", result1)

	//result2 := divMod(num1, num2)

	quotient, remainder := divMod(num1, num2)
	fmt.Println("quotient  =", quotient)
	fmt.Println("remainder  =", remainder)
}
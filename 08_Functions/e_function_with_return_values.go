package main

/*
Purpose: Function with two arguments and default return value

	NOTE:
	- return is the last statement in function execution
	-  When no return statement in function, cant assign and use as value

*/
import (
	"fmt"
)

func add(x int, y int) int {
	var output = x + y
	return output
}

func add1(x int, y int) (output int) {
	output = x + y
	// return keyword is mandatory when defined above
	//return output
	return
}

func divMod(x, y int) (q int, r int) {
	q = x / y
	r = x % y
	return q, r
}


func main(){
	num1 := 100
	num2 := 200

	result := add(num1, num2)
	fmt.Println("result     =", result)

	result1 := add1(num1, num2)
	fmt.Println("result1    =", result1)
}
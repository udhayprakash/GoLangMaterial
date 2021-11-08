package main

import "fmt"

/*
Purpose: Function with two arguments and default return value

	NOTE:
	- return is the last statement in function execution
	-  When no return statement in function, cant assign and use as value
*/

func add(x int, y int) int {
	var output = x + y
	return output
}

func add1(x, y int) (output int) {
	output = x + y // return keyword is mandatory when defined above
	//return output
	return      // <------ naked return
}

// Function to return Quotient and Remainder of Division
func divMod(x, y int)(q, r int){
	q = x / y
	r = x % y
	return q, r

}

func main() {
	// num1 := 100
	// num2 := 200

	num1, num2 := 100, 200 // Tuple assignment
	fmt.Printf("add(%3d, %3d) = %3d\n", num1, num2, add(num1, num2))
	fmt.Printf("add1(%3d, %3d) = %3d\n", num1, num2, add1(num1, num2))
	
	addResult := add1(num1, num2)
	fmt.Println("addResult:", addResult)

	// divResult := divMod(num1, num2)
	// assignment mismatch: 1 variable but divMod returns 2 values

	quotient, remainder := divMod(num1, num2)
	fmt.Println("quotient :", quotient)
	fmt.Println("remainder:", remainder)

	
	_, remainder1 := divMod(num1, num2)
	fmt.Println("remainder1:", remainder1)
}

 
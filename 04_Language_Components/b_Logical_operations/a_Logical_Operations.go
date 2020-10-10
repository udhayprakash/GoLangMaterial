package main

import "fmt"

// &&(AND), ||(OR), !(NOT)

func main(){
	// Logical AND Operation
	fmt.Println("true && true 	=", true && true)
	fmt.Println("true && false 	=", true && false)
	fmt.Println("false && true 	=", false && true)
	fmt.Println("false && false =", false && false)
	fmt.Println()


	// Logical OR Operation - true if atleast one among them is true
	fmt.Println("true || true 	=", true || true)
	fmt.Println("true || false 	=", true || false)
	fmt.Println("false || true 	=", false || true)
	fmt.Println("false || false =", false || false)
	fmt.Println()

	// Logical Negation
	fmt.Println("!true			=", !true)
	fmt.Println("!false			=", !false)
	fmt.Println()

	result := (true && false) || (false && true) || !(false && false)
	//             false      ||       false     || !false
	//             false      ||       false     || true   = true
	fmt.Println("result			=", result)

}

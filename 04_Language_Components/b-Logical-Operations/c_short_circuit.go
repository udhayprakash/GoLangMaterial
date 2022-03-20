package main

import "fmt"

// short-circuit logic
func main() {

	expr3 := 0 && 1 // 0
	fmt.Println("expr3:=", expr3)

	// expr4 := 3 && 9 // 9
	// fmt.Println("expr4:=", expr4)

	// expr4 := 3 && 9 && 13 // 13
	// fmt.Println("expr4:=", expr4)

	// expr4 := 3 && 13 && 9         // 9
	// fmt.Println("expr4:=", expr4) // && is resulting last element

	// expr5 := 0 || 1 // 1
	// fmt.Println("expr5:=", expr5)

	// expr6 := 3 || 9 // 3
	// fmt.Println("expr6:=", expr6)

	// expr6 := 3 || 9 || 13 // 3
	// fmt.Println("expr6:=", expr6)

	// expr6 := 13 || 9 || 3         // 13
	// fmt.Println("expr6:=", expr6) // || is resulting first element
	// fmt.Println()

	// // && - returns 0, if anyone is 0; else the last value
	// fmt.Println("{0 && 1 && 2 && 3 = }")  // 0
	// fmt.Println("{1 && 2 && 3 && 4 = }")  // 4
	// fmt.Println("{1 && 2 && 0 && 4 = }")  // 0
	// fmt.Println("{1 && 2 && 3      = }") // 3
	// fmt.Println()

	// // || - will take first boolean True value
	// fmt.Println("{0 || 1 || 2 || 3    = }") // 1
	// fmt.Println("{1 || 2 || 3 || 4    = }") // 1
	// fmt.Println("{4.4 || 2 || 3 || 4  = }") // 4.4

}

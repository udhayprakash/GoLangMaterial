package main

import "fmt"

// short-circuit logic
func main() {

	// expr3 := 0 && 1 // 0
	// fmt.Println("expr3:=", expr3)

	// num1 := 0
	// num2 := 1
	// expr3 := num1 && num2 // 0
	// fmt.Println("expr3:=", expr3)

	num1 := 23113

	cond1 := (num1 > 0)
	cond2 := (num1%2 == 0) // even number
	cond3 := (num1 < 10)

	// shortcircuiting - OR opertaion - checks for first true part
	fmt.Println("OR  op - ", cond1 || cond2 || cond3)
	//                        true

	// shortcircuiting - AND opertaion - checks for first false part
	fmt.Println("AND op - ", cond1 && cond2 && cond3)
	// 
	// NOTE: works only if there were same operation in entire expression
}

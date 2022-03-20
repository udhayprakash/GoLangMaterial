package main

import "fmt"

func main() {
	expr1 := (45 <= 45) || (3 > 333)
	//          true    ||    false   = true
	fmt.Println("expr1 = ", expr1)

	// NOTE: && has higher operator precendece than ||
	expr2 := (45 <= 45) || (3 > 333) && (9 == 9)
	//           true   ||   false   &&    true
	//           true   ||          false  = true
	fmt.Println("expr2 = ", expr2)

	// 23 < 73 < 45
	fmt.Println("23 < 73 && 73 < 45 =", 23 < 73 && 73 < 45)

	// 89 < 73 < 99 < 999 < 0
	fmt.Println("89 <73 && 73 <99 && 99< 999 &&  999< 0=", 89 < 73 && 73 < 99 && 99 < 999 && 999 < 0)
}

/*
NOTE:  Execution Flow: top-> bottom; left-> right

assignments
	1) ((45 <= 45) || (3 > 333)) && (9 == 9)
	2) (45 <= 45) || (3 > 333) || (9 == 9)
	3) !!true

*/

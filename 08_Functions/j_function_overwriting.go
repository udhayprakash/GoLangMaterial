package main

/*
Purpose: Function overwriting Problem

	NOTE: Redeclaration of the same function/variable is not permitted.

*/

import "fmt"

func myAddition(var1, var2, var3 int) int {
	return var1 + var2 + var3
}

//
//func myAddition(num1, num2 int)(int) {
//	return num1 + num2
//}

func main() {
	fmt.Println("myAddition(10, 20, 30):", myAddition(10, 20, 30))
	//fmt.Println("myAddition(10, 20)    :", myAddition(10, 20))

}

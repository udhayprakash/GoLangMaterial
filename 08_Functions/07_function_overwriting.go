package main
/*
Purpose: Function overwriting Problem

	NOTE: Redeclaration of the same function/variable is not permitted.
*/

import "fmt"

func addition(var1, var2, var3 int)(int) {
	return var1 + var2 + var3
}

//func addition(num1, num2 int)(int) {
//	return num1 + num2
//}


func main() {
	fmt.Println("addition(10, 20, 30):", addition(10, 20, 30))
	//fmt.Println("addition(10, 20)    :", addition(10, 20))

}

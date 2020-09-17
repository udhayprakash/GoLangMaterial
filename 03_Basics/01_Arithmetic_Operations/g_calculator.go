package main

import "fmt"

func main() {
	// Calculator Application

	//num1 := 100
	//num2 := 200

	var num1, num2 int32
	fmt.Println("Enter the values of num1 and num2:")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	fmt.Println("num1 + num2 = ", num1+num2)
	fmt.Println("num1 - num2 = ", num1-num2)
	fmt.Println("num1 * num2 = ", num1*num2)
	fmt.Println("num1 / num2 = ", num1/num2)

}

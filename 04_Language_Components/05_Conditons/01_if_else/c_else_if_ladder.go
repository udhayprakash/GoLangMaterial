package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Println("PLease enter two numbers, to check relationship:")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	if num1 > num2 {
		fmt.Println("num1 > num2")
	} else if num1 < num2 {
		fmt.Println("num1 < num2")
	} else { // num1 == num2
		fmt.Println("num1 == num2")
	}
}

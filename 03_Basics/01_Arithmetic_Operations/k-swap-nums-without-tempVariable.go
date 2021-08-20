package main

import (
	"fmt"
)

func main() {
	num1, num2 := 100, 200
	
	// Method 1 - swapping using 3rd variable
	var num3 int
	fmt.Printf("Before Swap, num1 = %d, num2 = %d\n", num1, num2)
	num3 = num1
	num1 = num2
	num2 = num3
	fmt.Printf("After Swap, num1 = %d, num2 = %d\n", num1, num2)
	
	// Method2 - swapping without 3rd variable
	fmt.Printf("Before Swap, num1 = %d, num2 = %d\n", num1, num2)
	num1, num2 = 100, 200
	num1 = num2 - num1
	num2 = num1 + num2
	num1 = num2 - num1
	fmt.Printf("After Swap, num1 = %d, num2 = %d\n", num1, num2)
	
	// Method 3 - Direct swap 
	fmt.Printf("Before Swap, num1 = %d, num2 = %d\n", num1, num2)
	num1, num2 = 100, 200
	num1, num2 = num2, num1 
	fmt.Printf("After Swap, num1 = %d, num2 = %d\n", num1, num2)

}

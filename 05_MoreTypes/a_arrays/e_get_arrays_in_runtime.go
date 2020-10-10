package main

import "fmt"

func main(){
	var numbers [10]int

	// Method 1
	//fmt.Println("Enter 10 numbers:")
	//fmt.Scanf("%d", &numbers)
	//fmt.Println("numbers=", numbers)

	// Method 2
	for index:= 0; index < len(numbers); index++{
		fmt.Printf("\nEnter %d position value:", index)
		fmt.Scanf("%d", &numbers[index])
	}
	fmt.Println("numbers=", numbers)
}
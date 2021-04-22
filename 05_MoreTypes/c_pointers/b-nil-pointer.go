package main 


import (
	"fmt"
)

func main(){
	var ptr *float64
	fmt.Printf("The value of ptr is : %x\n", ptr)
	fmt.Println("Before ptr != nil :", ptr != nil)

	num2 := 123.213

	ptr = &num2
	fmt.Println("After  ptr != nil :", ptr != nil)

	// myStr := "Golang"
	// ptr = &myStr

}

// NOTE: You are refer values of same data types as defined in pointer
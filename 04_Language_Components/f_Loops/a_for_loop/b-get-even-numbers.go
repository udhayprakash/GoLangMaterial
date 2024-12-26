package main

import (
	"fmt"
)

func main() {
	fmt.Println("Even number between 56 & 100:")

	for num := 56; num <= 100; num++ {
		if num%2 == 0 {
			// fmt.Println(num)
			fmt.Print(num, " ")
		}
	}

	fmt.Println()

	for num := 56; num <= 100; num++ {
		if num%2 == 0 {
			fmt.Print(num, " ")
		}
	}

	// fmt.Println("num = ", num)
	// undefined: num
}

//  golang supports only ->  i++,  i--
//  doesnt support       ->  ++i, --i


// variable defined in for loop definition , cant be accessed outside 

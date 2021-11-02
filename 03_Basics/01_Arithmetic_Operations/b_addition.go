package main

import "fmt"

func main() {
	fmt.Println("1 + 1 		=", 1+1)
	fmt.Println("1.0 + 1.0 	=", 1.0+1.0)
	fmt.Println("1.0 + 1.1 	=", 1.0+1.1)

	fmt.Println("1.3 + 222  =", 1.3+222)

	// fmt.Println(12 + true) // error: incompatible types in binary expression

	fmt.Println("'a' + 10   =", 'a'+10) // 107
}
 
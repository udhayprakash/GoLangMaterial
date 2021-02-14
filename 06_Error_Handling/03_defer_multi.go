package main

import "fmt"

func main() {
	//  Deferred functions may read and assign to 
	// the returning function's named return values.
	fmt.Println("c() =", c())  // 2
}


func c() (i int) {
	defer func() { i++ }()  // i = 2
	return 1      		// i = 1
}
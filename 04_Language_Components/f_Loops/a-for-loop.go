package main

import "fmt"

/*
for initialization; condition; post {
	// zero or more statements
	}

*/

func main(){
	// Method 1
	var i int 
	for i = 0; i <= 5; i++{
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Method 2 
	for j := 0; j <=5; j++{
		fmt.Printf("%d\t", j)
	}
	fmt.Println()

	// Method 3  --> like while loop style
	j := 0
	for j <= 5{
		fmt.Printf("%d\t", j)
		j++
	}

}
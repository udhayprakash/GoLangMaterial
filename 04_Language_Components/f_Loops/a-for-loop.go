package main

import "fmt"

/*
for initialization; condition; post {
	// zero or more statements
	}

*/

func main() {
	// Method 1
	var i int
	for i = 0; i <= 5; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Method 2
	for j := 0; j <= 5; j++ {
		fmt.Printf("%d\t", j)
	}
	fmt.Println()

	// Method 3
	j := 0
	for j <= 5 {
		fmt.Printf("%d\t", j)
		j++
	}
	fmt.Println()

}

/*
NOTE:
	1. All three (initialization,condition and post) are optional.
	2. prefix unary operations (--i, ++i) are illegal
	3. i++ is same as i += 1, which is same as i = i + 1
	   i-- is same as i -= 1, which is same as i = i - 1
	4. i++ and i-- are treated as statements, not expressions.
	   It means, j = i++ is illegal in Go unlike C language.
*/

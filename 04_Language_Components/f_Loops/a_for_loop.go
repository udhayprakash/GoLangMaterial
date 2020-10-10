package main

import "fmt"

/*
for initialization; condition; post {
	// zero or more statements
	}

 */
func main(){
	// Method 1
	for i:= 1; i <= 10; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Method 2
	j := 1
	for j <= 10 {
		fmt.Printf("%d\t", j)
		j++
	}
}

/*
NOTE:
	1. All three (initialization,condition and post) are optional.
	2. prefix unary operations (--i, ++i) are illegal
	3. i++ is same as i += 1, which is same as i = i + 1
	   i-- is same as i -= 1, which is same as i = i - 1
	4. i++ and i-- are treated as statements, not expressions.
	   It means, j = i++ is illegal in Go unlike C language. */

package main

import "fmt"

func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 10 {  // - like while loop in C
		fmt.Println(i)
		i++ //i += 1  // i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

/*
for initialization; condition; post {
	// zero or more statements
	}

NOTE:
	1. All three (initialization,condition and post) are optional.
	2. prefix unary operations (--i, ++i) are illegal
	3. i++ is same as i += 1, which is same as i = i + 1
	   i-- is same as i -= 1, which is same as i = i - 1
	4. i++ and i-- are treated as statements, not expressions.
	   It means, j = i++ is illegal in Go unlike C language.

*/

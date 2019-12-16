package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ //i += 1  // i = i + 1
	}
}

// all unary operations (i++, i--, --i, ++i) are not supported

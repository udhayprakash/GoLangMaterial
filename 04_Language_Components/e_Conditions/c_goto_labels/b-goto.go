package main

import "fmt"

func main() {

	i := 0

	goto Here

Here:
	fmt.Println("i =", i)
	i++
	if i <= 10 {
		goto Here
	}
}

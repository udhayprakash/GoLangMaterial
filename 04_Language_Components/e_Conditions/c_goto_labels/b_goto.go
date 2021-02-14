package main

import "fmt"

func main() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i <= 10 {
		goto Here
	}
}

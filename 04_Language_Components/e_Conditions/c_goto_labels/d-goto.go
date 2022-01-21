package main

import "fmt"

func main() {
	var a int = 1

Label:
	for a < 5 {
		if a == 3 {
			a = a + 1
			goto Label
		}
		fmt.Printf("Number is :%d\n", a)
		a++
	}
}

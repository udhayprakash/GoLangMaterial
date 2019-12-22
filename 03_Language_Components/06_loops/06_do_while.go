package main

import "fmt"

/*
Go Doesnt support do-while loop
*/
func hello()  {
	fmt.Println("Hello world")
}
func main() {
	for ok := true; ok; ok = condition {
		hello()
	}
}

package main

import (
	. "fmt"
)

/*
The dot operator means you can omit the package name
when you call functions inside of that package.
*/
func main() {
	Println("Dot operator based import usage") // fmt.Println
	println("Dot operator based import usage") // println
}

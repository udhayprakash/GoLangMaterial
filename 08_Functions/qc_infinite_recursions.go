package main

import "fmt"

func factorial(num int) int {
	//if num == 0{
	//	return  1
	//}
	return num * factorial(num - 1)
}

func main(){
	fmt.Println("factorial(9) =", factorial(9))
	//runtime: goroutine stack exceeds 1000000000-byte limit
	//fatal error: stack overflow
}
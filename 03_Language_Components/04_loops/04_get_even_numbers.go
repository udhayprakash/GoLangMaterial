package main

import "fmt"


func main(){

	fmt.Println("Even no.s between 0 & 10")
	for i:= 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}

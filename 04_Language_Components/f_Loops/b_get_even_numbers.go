package main

import "fmt"

func main() {
	fmt.Println("Even number between 0 & 25:")
	for num := 0; num <= 25; num++ {
		if num%2 == 0 {
			//fmt.Println(num%2 == 0, num)
			fmt.Printf("%d ", num)
		}
	}
}

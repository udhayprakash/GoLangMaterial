package main

import "fmt"

func main() {
	fmt.Printf("isPowerOf2(%4d) = %v\n", 34, isPowerOf2(34))
	fmt.Printf("isPowerOf2(%4d) = %v\n", 32, isPowerOf2(32))
	fmt.Printf("isPowerOf2(%4d) = %v\n", 1024, isPowerOf2(1024))
}

func isPowerOf2(num int) bool {
	for num != 1 {
		if num%2 != 0 {
			return false
		}
		num = num / 2
	}
	return true
}

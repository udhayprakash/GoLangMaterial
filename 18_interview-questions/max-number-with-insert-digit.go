package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(268) == 5268)
	fmt.Println(Solution(670) == 6750)
	fmt.Println(Solution(0) == 50)
	fmt.Println(Solution(-999) == -5999)
}

func Solution(N int) int {
	NumStr := strconv.Itoa(N)
	tempStr := ""
	result := 0
	for index := range NumStr {
		tempStr = NumStr[:index] + "5" + NumStr[index:]
		x, _ := strconv.Atoi(tempStr)
		if result == 0 || x > result {
			result = x
		}

	}
	return result
}

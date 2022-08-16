package main

import "fmt"

func main() {
	var inputNum, result int

	for i := 0; i < 5; i++ {

		fmt.Scanf("%d", &inputNum)
		result = result*10 + inputNum
	}
	fmt.Println("result:", result)

}

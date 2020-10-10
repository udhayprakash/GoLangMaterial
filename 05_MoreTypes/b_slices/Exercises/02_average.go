package main

import "fmt"

func main() {
	values := []float32{78, 33, 2, 34, 22, 66}

	var result float32
	for _, value := range values {
		result += value
	}
	//fmt.Println("Average:", result/len(values)) // 	fmt.Println("Average:", result/len(values))
	fmt.Println("Average:", result/float32(len(values)))
}

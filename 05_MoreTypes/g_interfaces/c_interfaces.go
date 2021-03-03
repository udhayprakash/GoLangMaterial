package main

import "fmt"

func ReturnData() interface{} {
	return "asd"
	return 5
}

func main() {
	result := ReturnData()
	fmt.Println("result:", result)
}

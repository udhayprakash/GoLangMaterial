package main

import (
	"fmt"
)

func getNum(num int) int {
	return num
}

// func getAddressLocation(num int) int {
// 	fmt.Println(reflect.TypeOf(&num))
// 	return &num
// }

func getAddressLocation(num int) *int {
	// fmt.Println(reflect.TypeOf(&num))
	return &num
}

// Generic function
func getAddressLocationGeneric(variable interface{}) *interface{} {
	return &variable
}

func main() {
	num1 := 123                                               // int
	fmt.Println(num1, getNum(num1), getAddressLocation(num1)) // 123 123 0xc0000120a8

	name := "Deepa"
	fmt.Println(name, getAddressLocationGeneric(name)) // 123 123 0xc0000120a8

	result := true //  bool
	fmt.Println(result, getAddressLocationGeneric(result))

	myArray := [...]int{11, 22, 33} // array
	fmt.Println(myArray, getAddressLocationGeneric(myArray))

	fmt.Println("myArray[0]=", myArray[0], &myArray[0])
	fmt.Println("myArray[1]=", myArray[1], &myArray[1])
	fmt.Println("myArray[2]=", myArray[2], &myArray[2])

}

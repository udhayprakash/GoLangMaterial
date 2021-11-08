package main

import "fmt"

// func getAddressLocation(num int) int {
// 	return &num
// }
// cannot use &num (type *int) as type int in return argument

// It is perfectly safe for a function to return
// the address of a local variable.
func getAddressLocation(num int) *int {
	return &num
}

// Generic Function for all data types
func getAddressLocation3(num interface{}) *interface{} {
	return &num
}

func main() {
	num1 := 123 // int
	fmt.Println(num1, getAddressLocation(num1))

	name := "udhay" // string
	fmt.Println(name, getAddressLocation3(name))

	result := true //  bool
	fmt.Println(result, getAddressLocation3(result))

	myArray := [...]int{11, 22, 33} // array
	fmt.Println(myArray, getAddressLocation3(myArray))

	fmt.Println("myArray[0]=", myArray[0], &myArray[0])
	fmt.Println("myArray[1]=", myArray[1], &myArray[1])
	fmt.Println("myArray[2]=", myArray[2], &myArray[2])

}
 
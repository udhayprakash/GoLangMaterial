package main

import (
	"fmt"
)

// `Type switch`.
// func main() {
// 	fmt.Print("Enter some number:")

// 	var value int
// 	fmt.Scanf("%d", &value)
// 	fmt.Println(value, reflect.TypeOf(value))

// 	var valu2 float32
// 	fmt.Scanf("%f", &valu2)
// 	fmt.Println(valu2, reflect.TypeOf(valu2))

// 	var valu3 interface{}
// 	fmt.Scanf("%v", &valu3)
// 	fmt.Println(valu3, reflect.TypeOf(valu3))

// }

func main() {
	getDataType(true)
	getDataType(1)
	getDataType("Hello")
	getDataType('h')
	getDataType(123.21)
}

// It can be used to discover the type of an interface
// value.
func getDataType(value interface{}) {
	fmt.Printf("\nvalue =%v \tdataType - %T\n", value, value)

	// fmt.Println("value.(type)=", value.(type))
	// use of .(type) outside type switch

	switch t:= value.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int, int32, int64:
		fmt.Println("I'm an int")
	case float32, float64:
		fmt.Println("I'm a float")
	case string:
		fmt.Println("I'm a string")
	default:
		fmt.Printf("Don't know type %T\n", t)

	}
}
 
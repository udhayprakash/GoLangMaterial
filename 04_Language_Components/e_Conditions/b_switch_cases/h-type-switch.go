package main

import "fmt"

// A type `switch` compares types instead of values.
// It can be used to discover the type of an interface
// value.
func main() {
	getDataType(true)
	getDataType(1)
	getDataType("Hello")
	getDataType('h')
	getDataType(123.21)
}

func getDataType(value interface{}) {
	fmt.Printf("\n\nvalue =%v \tdataType - %T\n", value, value)

	// fmt.Println("value.(type)=", value.(type))
	// use of .(type) outside type switch

	switch t := value.(type) {
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

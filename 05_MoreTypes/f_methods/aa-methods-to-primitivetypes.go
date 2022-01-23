package main

import "fmt"

/*
	Methods can be created for any data type.

	we can declare method in a different package, to that of type.
	Thats why we can't add a method to predefined names types.

	so, aliase types need to be created, for premitive types, for
	using as methods.

SYNTAX:

	func (receiverName receiverType) methodName(param1 paramType) (returnType1) {
		…
	}

*/
type Number int
type FloatNum float32
type MyStr string

// type MyInterface interface{}

// method
func (num Number) displayText() {
	fmt.Printf("number = %d\n", num)
}

func (num FloatNum) displayTextFloat() {
	fmt.Printf("number = %f\n", num)
}

func (str MyStr) displayTextString() {
	fmt.Printf("string = %s\n", str)
}

// func (gen MyInterface) displayTextMyInterface() {
// 	fmt.Printf("value = %v  type = &[1]T\n", gen)
// } // invalid receiver type MyInterface (MyInterface is an interface type)

func main() {
	// var num1 int = 123

	var num1 Number = 123
	num1.displayText()

	// num2 := 123.3
	var num2 FloatNum = 123
	num2.displayTextFloat()

	var mystr1 MyStr = "Go Language"
	mystr1.displayTextString()

	// var value1 MyInterface = 123
	// value1.displayTextMyInterface()

	// -----------------
	v1 := Number(5)
	v2 := Number(6)
	fmt.Println("v1.Add(v2)      =", v1.Add(v2))
	fmt.Println("v1.Add(v2) + 12 =", v1.Add(v2)+12)

	// -----------
	fmt.Println(mystr1.Required())         // true
	fmt.Println(MyStr("Hello").Required()) // true
}

func (i Number) Add(j Number) Number {
	return i + j
}

func (s MyStr) Required() bool {
	return s != ""
}

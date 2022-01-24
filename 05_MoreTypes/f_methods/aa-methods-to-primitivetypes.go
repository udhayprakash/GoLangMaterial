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
type MyInt int
type MyFloat float64
type MyStr string

// methods
func (num MyInt) displayText() {
	fmt.Printf("MyInt   = %d\n", num)
}

func (num MyFloat) displayText() {
	fmt.Printf("MyFloat = %f\n", num)
}

func (str MyStr) displayText() {
	fmt.Printf("MyStr   = %s\n", str)
}

func (i MyInt) Add(j MyInt) MyInt {
	return i + j
}

func (s MyStr) IsNotEmpty() bool {
	return s != ""
}

func main() {
	// var num1 int = 123

	var num1 MyInt = 123
	num1.displayText()

	// num2 := 123.3
	var num2 MyFloat = 123.3
	num2.displayText()

	var mystr1 MyStr = "Go Language"
	mystr1.displayText()

	//----------------
	v1 := MyInt(5)
	v2 := MyInt(6)

	fmt.Println("v1.Add(v2)      =", v1.Add(v2))    // 11
	fmt.Println("v1.Add(v2) + 12 =", v1.Add(v2)+12) // 23

	fmt.Println("v1.Add(22)      =", v1.Add(22))    // 27
	fmt.Println("v1.Add(22) + 12 =", v1.Add(22)+12) // 39

	fmt.Println("mystr1.IsNotEmpty() =", mystr1.IsNotEmpty()) // true
	// fmt.Println( "Ravi".IsNotEmpty())
	// "Ravi".IsNotEmpty undefined (type string has no field or method IsNotEmpty)

	fmt.Println(MyStr("Ravi").IsNotEmpty()) // true
	fmt.Println(MyStr("").IsNotEmpty())     // false
}

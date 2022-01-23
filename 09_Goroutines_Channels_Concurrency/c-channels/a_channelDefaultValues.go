package main

/*
Purpose: Channels
	- are the pipes that connect concurrent goroutines.
	- You can send values into channels from one goroutine
      and receive those values into another goroutine.

*/
import "fmt"

// There are TWO ways of creating Channels
func main() {
	// var myINt int
	var myIntChannel chan int
	fmt.Printf("myIntChannel   - value: %v \n", myIntChannel) // nil
	fmt.Printf("myIntChannel   - type : %T \n", myIntChannel) // chan int

	var myFloatChannel chan float64
	fmt.Printf("\nmyFloatChannel - value: %v \n", myFloatChannel) // <nil>
	fmt.Printf("myFloatChannel - type : %T \n", myFloatChannel)

	var myBoolChannel chan bool
	fmt.Printf("\nmyBoolChannel - value: %v \n", myBoolChannel) // <nil>
	fmt.Printf("myBoolChannel - type : %T \n", myBoolChannel)

	//const myIntConstChannel chan
	// ERROR: const declaration cannot have type without expression

	var myStringChannel1 chan string
	fmt.Printf("\nmyStringChannel1 - value: %v \n", myStringChannel1) // nil
	fmt.Printf("myStringChannel1 - type : %T \n", myStringChannel1)

	myStringChannel := make(chan string)
	fmt.Printf("\nmyStringChannel - value: %v \n", myStringChannel) // 0xc00001a1e0
	fmt.Printf("myStringChannel - type : %T \n", myStringChannel)

	myInterfaceChannel := make(chan interface{})
	fmt.Printf("\nmyInterfaceChannel - value: %v \n", myInterfaceChannel) // 0xc00001a240
	fmt.Printf("myInterfaceChannel - type : %T \n", myInterfaceChannel)
}

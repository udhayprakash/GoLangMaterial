package main

/*
Purpose: Channels
	- are the pipes that connect concurrent goroutines.
	- You can send values into channels from one goroutine
      and receive those values into another goroutine.

*/
import "fmt"

func main() {
	var myIntChannel chan int
	fmt.Printf("myIntChannel   - value: %v \n", myIntChannel)
	fmt.Printf("myIntChannel   - type : %T \n", myIntChannel)

	var myFloatChannel chan float64
	fmt.Printf("\nmyFloatChannel - value: %v \n", myFloatChannel)
	fmt.Printf("myFloatChannel - type : %T \n", myFloatChannel)

	var myBoolChannel chan bool
	fmt.Printf("\nmyBoolChannel - value: %v \n", myBoolChannel)
	fmt.Printf("myBoolChannel - type : %T \n", myBoolChannel)

	//const myIntConstChannel chan
	// ERROR: const declaration cannot have type without expression

	var myStringChannel1 chan string
	fmt.Printf("\nmyStringChannel1 - value: %v \n", myStringChannel1)
	fmt.Printf("myStringChannel1 - type : %T \n", myStringChannel1)

	myStringChannel := make(chan string)
	fmt.Printf("\nmyStringChannel - value: %v \n", myStringChannel)
	fmt.Printf("myStringChannel - type : %T \n", myStringChannel)

	myInterfaceChannel := make(chan interface{})
	fmt.Printf("\nmyInterfaceChannel - value: %v \n", myInterfaceChannel)
	fmt.Printf("myInterfaceChannel - type : %T \n", myInterfaceChannel)

}

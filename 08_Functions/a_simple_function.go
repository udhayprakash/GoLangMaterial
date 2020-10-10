package main

import (
	"fmt"
	"reflect"
)

/*
Purpose: Functions
            - code re-usability
            - To modularize the problem
            - Better maintenance of the code
			- treated as first class objects
			- function definitions need not be before calls
Go requires explicit returns, i.e. it won’t automatically
return the value of the last expression.
*/

func myFunc() {
	//Function with no arguments and no return value
	fmt.Println("Function myFunc is called!!!")
}


func main(){
	fmt.Println("myFunc                ", myFunc)
	fmt.Println("reflect.TypeOf(myFunc)", reflect.TypeOf(myFunc))

	myFunc()
}

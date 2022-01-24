package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element // slice of interfaces

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 5)
	fmt.Println("Initially, list =", list)

	list[0] = 1                    //an int
	list[1] = "Hello"              //a string
	list[2] = Person{"Dennis", 70} // struct instance
	list[3] = true                 // bool
	list[4] = []int{11, 22, 33}    // slice
	fmt.Println("after initial, list =", list)

	for index, element := range list {
		fmt.Println(index, element)
		switch value := element.(type) { //  Type Switch
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		case bool:
			fmt.Printf("list[%d] is a boolean and its value is %v\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

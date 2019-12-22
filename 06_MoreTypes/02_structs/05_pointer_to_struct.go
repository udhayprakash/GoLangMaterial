package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	RollNumber int
	Name       string
}

func main() {
	// instance of student struct type
	s1 := Student{14, "Rahul Dravid"}
	fmt.Println("\n", reflect.TypeOf(s1), s1)

	// Pointer to the student struct
	p_s1 := &s1
	fmt.Println("\n", reflect.TypeOf(p_s1), p_s1)

	// Accessing struct fields via pointer
	fmt.Println("\n", (*p_s1).Name, s1.Name)
	fmt.Println("(*p_s1).Name == s1.Name :", (*p_s1).Name == s1.Name)

	// Updating struct value
	fmt.Println("\nBefore s1.RollNumber=", s1.RollNumber)
	s1.RollNumber = 23
	fmt.Println("After    s1.RollNumber=", s1.RollNumber)

	// updating value via pointer reference
	p_s1.Name = "KL Rahul"
	fmt.Println("\n", (*p_s1).Name, s1.Name)

}

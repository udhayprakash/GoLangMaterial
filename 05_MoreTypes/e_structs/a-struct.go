package main

import (
	"fmt"
	"reflect"
)

/*
Purpose: Structs
	- Go’s structs are typed collections of fields.
	- They’re useful for grouping data together to form
      records.
*/
type Person struct {
	FirstName string // each value is called field
	LastName  string
	Age       int
	Salary    float64
}

func main() {
	// Method 1 : First declaration and then initialization
	var num1 int
	var p Person
	fmt.Printf("num1= %7d - > type -%[1]T - detailed-%v\n", num1, reflect.TypeOf(num1).Kind())
	fmt.Printf("p   = %v - > type -%[1]T - detailed-%v\n", p, reflect.TypeOf(p).Kind())

	// dot operation
	fmt.Printf(`Zero-values
		p.FirstName = %s     
		p.LastName = %s      
		p.Age = %d 
		p.Salary = %f`, p.FirstName, p.LastName, p.Age, p.Salary)

	// assigning values to struct fields
	p.FirstName = "Narendra"
	p.LastName = "Modi"
	p.Age = 75
	p.Salary = 100000

	fmt.Println()
	fmt.Printf(`After initialization
		p.FirstName = %s     
		p.LastName = %s      
		p.Age = %d
		p.Salary = %f`, p.FirstName, p.LastName, p.Age, p.Salary)

	// =================================================
	// Method 2: Declaration & Initialization together
	// var p1 = Person{"Saurav"} // Compiler Error: too few values in struct initializer
	var p1 = Person{"Saurav", "Ganguly", 43, 123213}

	fmt.Println()
	fmt.Printf(`Zero-values
		p1.FirstName = %s
		p1.LastName = %s
		p1.Age = %d
		p1.Salary = %d`, p1.FirstName, p1.LastName, p1.Age, p1.Salary)

	// Method 3: Naming fields while initializing struct
	var p2 = Person{FirstName: "Alien"} // LastName: "", Age: 0
	// p2 = Person{}                       // FirstName: "", LastName: "", Age: 0

	p2 = Person{
		FirstName: "Sachin",
		LastName:  "Tendulkar",
		Age:       43,
	}
	fmt.Println()
	fmt.Printf(`Zero-values
		p2.FirstName = %s
		p2.LastName = %s
		p2.Age = %d
		p2.Salary = %f`, p2.FirstName, p2.LastName, p2.Age, p2.Salary)

}

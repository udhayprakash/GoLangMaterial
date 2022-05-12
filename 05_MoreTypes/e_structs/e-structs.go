package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Voter struct {
		name           string
		age            int
		eligibleToVote bool
	}
	fmt.Println("Voter{}:", Voter{})
	fmt.Println(Voter{"udhay", 23, true})
	fmt.Println(Voter{name: "udhay", age: 23, eligibleToVote: true})
	fmt.Println()

	// When giving one value
	// fmt.Println(Voter{"udhay"})
	fmt.Println(Voter{name: "udhay"})

	// pointer ops in structs
	fmt.Println("&Voter{name:\"udhay\"}=", &Voter{name: "udhay"}) // &{udhay 0 false}

	ptr := &Voter{name: "udhay"}
	fmt.Println("ptr:", ptr)         // &{udhay 0 false}
	fmt.Println("*ptr:", *ptr)       // {udhay 0 false}
	fmt.Println(reflect.TypeOf(ptr)) // *main.Voter
	fmt.Println()

	s := Voter{name: "Sean", age: 50}
	fmt.Println(s.name) // Sean
	s.name = "Ravi"     // updating with struct instance
	fmt.Println(s.name) // Ravi

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51 // updating with struct instance pointer
	fmt.Println(sp.age)

	fmt.Println("s =", s) // {Ravi 51 false}
}

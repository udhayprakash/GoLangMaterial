package main

import "fmt"

type person struct {
	name string
	age  int
	eligibleToVote bool
}

func NewPerson(name string) *person {
	// Function returning `person` struct type
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20, true})

	fmt.Println(person{name: "Alice", age: 30, eligibleToVote:true})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(NewPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

package main

import "fmt"

type person struct {
	name           string
	age            int
	eligibleToVote bool
}

func main() {
	fmt.Println(person{"Udhay", 28, true})

	fmt.Println(person{name: "Udhay", age: 28, eligibleToVote: true})
	fmt.Println(person{age: 28, eligibleToVote: true, name: "Udhay"})
	fmt.Println()

	//fmt.Println(person{"Udhay"}) // too few values in person literal
	fmt.Println(person{name: "Udhay"})
	fmt.Println()

	fmt.Println(&person{name: "Udhay"})
	fmt.Println(NewPerson("Udhay"))
	fmt.Println()

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}

func NewPerson(name string) *person {
	// Function returning `person` struct type
	p := person{name: name}
	p.age = 23
	return &p
}

package main

import (
	"encoding/json"
"fmt"
)
type Person struct {
	Name   string // 1st letter cap is public, while small cap is private
	email  string
	Gender string
	Age    int
}

func (p *Person) SetEmail(email string) {
	p.email = email
}

func (p Person) Email() string {
	return p.email
}

var v = []byte(`[
     {
     "name": "Denis Silva Costa",
     "email": "d*@gmail.com",
     "gender": "Male",
     "age": 27
     },
     {
     "name": "Ariana Cursino",
     "email": "a@a.com",
     "gender": "Female",
     "age": 31
     }
     ]`)

func main() {

	var people []Person
	err := json.Unmarshal(v, &people)
	if err != nil {
	}

	// email WILL NOT be displayed because it is private
	fmt.Println(people)

	// set data to private variable via SetEmail method

	people[0].SetEmail("d@gmail.com")
	people[1].SetEmail("a@a.com")

	// Retrieve data from private variables via Email method
	fmt.Println(people[0].Email())
	fmt.Println(people[1].Email())

}

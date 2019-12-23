package main

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	return
}
type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	// Map
	myMap := map[string]string{
		"a": "apple",
		"b": "ball",
		"c": "cat",
	}
	fmt.Println(myMap)
	PrettyPrint(myMap)
	fmt.Println()

	// Slice
	slcD := []string{"apple", "peach", "pear"}
	fmt.Println(slcD)
	PrettyPrint(slcD)
	fmt.Println()

	// struct
	p1 := &Person{"Sam", 20, []string{"cricket", "football"}}
	fmt.Println(p1)
	PrettyPrint(p1)
}

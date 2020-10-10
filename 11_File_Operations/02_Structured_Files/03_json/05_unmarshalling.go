package main

import "encoding/json"
import "fmt"

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	var bob, alice User
	json.Unmarshal([]byte(`{"name":"bob","age":25}`), &bob)
	err := json.Unmarshal([]byte(`{"name":"alice","address":"NY"}`), &alice)

	fmt.Printf("err: %v\n", err)
	fmt.Printf("bob: %v\n", bob)
	fmt.Printf("alice: %v\n", alice)
}
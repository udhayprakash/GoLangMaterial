package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Nationality string  `json:"nationality"`
	Address     Address `json:"address"`
}
type Address struct {
	City  string
	State string
}

func main() {
	userProfile := Profile{Name: "Sibi", Age: 27, Nationality: "Indian", Address: Address{City: "Bangalore", State: "Karnataka"}}
	fmt.Printf("%+v", userProfile)

	byteArray, err := json.Marshal(userProfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	jsondata := []byte(`{"Name":"Sibi","Age":27,"Nationality":"Indian","Address":{"City":"Bangalore","State":"Karnataka"}}`)
	var f map[string]interface{}
	err = json.Unmarshal(jsondata, &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The decoded json data : ", f)
}

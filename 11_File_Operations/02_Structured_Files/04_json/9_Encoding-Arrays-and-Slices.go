package main

import (
	"encoding/json"
	"fmt"
)


type Bird struct {
	Species     string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

func main(){
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// Now we pass a slice of two pigeons
	data, _ := json.Marshal([]*Bird{pigeon, pigeon})
	fmt.Println(string(data))
}

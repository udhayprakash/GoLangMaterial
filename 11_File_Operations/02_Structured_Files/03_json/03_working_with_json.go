package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{
		  "species": "pigeon",
		  "decription": "likes to perch on rocks"
		}`
	type Bird struct {
		Species     string
		Description string
	}
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s\n",
		bird.Species, bird.Description)
	// Attribute names will  be of same case as in json.
	// So, Description is empty string in output

}

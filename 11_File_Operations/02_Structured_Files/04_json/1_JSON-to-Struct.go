package main

import (
	"encoding/json"
	"fmt"
)

// Online JSON-to-struct tool - https://mholt.github.io/json-to-go/
type Bird struct {
	Species     string
	Description string
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
	//Species: pigeon, Description: likes to perch on rocks
}

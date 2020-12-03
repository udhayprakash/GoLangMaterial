package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// The keys need to be strings, the values can be
	// any serializable value
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
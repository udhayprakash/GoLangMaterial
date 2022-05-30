package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Users  []string
	Groups []string
}

func main() {
	file, _ := os.Open("a_conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users) // output: [UserA, UserB]
}

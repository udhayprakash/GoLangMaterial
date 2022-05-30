package main

import (
	"encoding/json"	"fmt"
)


 func main() {
 	// Given a possibly complex JSON object
 	msg := "{\"assets\" : {\"old\" : 123}}"

 	// We only know our top-level keys are strings
 	mp := make(map[string]interface{})

 	// Decode JSON into our map
 	err := json.Unmarshal([]byte(msg), &mp)
 	if err != nil {
 		println(err)
 		return
 	}

 	// See what the map has now
 	fmt.Printf("mp is now: %+v\n", mp)

 	// Iterate the map and print out the elements one by one
 	// Note: that mp has to be deferenced here or range will fail
 	for key, value := range mp {
 		fmt.Println("key:", key, "value:", value)
 	}
 }
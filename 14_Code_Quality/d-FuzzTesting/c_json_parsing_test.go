package main

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ParseJSON(data []byte) (*Person, error) {
	var p Person
	err := json.Unmarshal(data, &p)
	return &p, err
}

func FuzzParseJSON(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := ParseJSON(data)
		if err != nil {
			// Ignore invalid JSON inputs
			return
		}
	})
}

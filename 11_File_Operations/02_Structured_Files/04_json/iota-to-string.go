package main

import (
	"encoding/json"
	"fmt"
)

const (
	GenderNotSet = iota
	GenderMale
	GenderFemale
	GenderOther
)

type Human struct {
	Gender int `json:"gender"`
}

func main() {
	me := Human{
		Gender: GenderMale,
	}
	prettyJSON, _ := json.MarshalIndent(me, "", "    ")
	fmt.Println(string(prettyJSON))
}

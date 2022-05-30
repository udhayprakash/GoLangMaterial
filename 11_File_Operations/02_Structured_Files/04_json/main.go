package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Gender Gender `json:"gender"`
}

func main() {
	me := Human{
		Gender: GenderMale,
	}
	prettyJSON, _ := json.MarshalIndent(me, "", "    ")
	fmt.Println(string(prettyJSON))
}

package main

import (
	"encoding/json"
	"fmt"
)

/*
{
  "species": "pigeon",
  "decription": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}

*/

type Dimensions struct {
	Height int
	Width int
}
type Bird struct {
	Species string
	Description string
	Dimensions Dimensions
}

func main(){
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var birds Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf(bird)
	// {pigeon likes to perch on rocks {24 10}}
}

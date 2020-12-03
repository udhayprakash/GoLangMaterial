package main

import "encoding/json"

func main()  {
	myJsonString := `{"some":"json"}`
	json.Unmarshal([]byte(myJsonString), &myStoredVariable)

}
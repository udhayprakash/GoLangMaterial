package main
import (
  "encoding/json"
  "fmt"
)

// Struct defined
type ShotExample struct {
  Name string
  Shot string
}

func main() {

  // Converting Integer to JSON
  intJson, err := json.Marshal(22)
  if err == nil { 
    fmt.Println(string(intJson))
  } else {
    panic(err)
  }

  // Converting String to JSON
  strJson, err := json.Marshal("Hello World")
  if err == nil { 
    fmt.Println(string(strJson))
  } else {
    panic(err)
  }

  // Declaring a map
  mapVar := map[string]int{"num1": 1, "num2": 2}
  // Converting map to JSON
  mapJson, err := json.Marshal(mapVar)
  if err == nil { 
    fmt.Println(string(mapJson))
  } else {
    panic(err)
  }

  // Declaring struct object
  structVar := ShotExample{Name: "Hello World", Shot: "JSON"}
  // Converting map to JSON
  structJson, err := json.Marshal(structVar)
  if err == nil { 
    fmt.Println(string(structJson))
  } else {
    panic(err)
  }
}
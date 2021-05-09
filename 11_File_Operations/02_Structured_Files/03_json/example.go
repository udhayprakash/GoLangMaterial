package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `
   {
      "server": {"host": "127.0.0.1", "path": "/some/silly/path", "port": 8080}
   }
   `
	var config struct {
		Server struct {
			Host, Path string
			Port       int
		}
	}
	json.Unmarshal([]byte(s), &config)
	fmt.Printf("%+v\n", config)
}

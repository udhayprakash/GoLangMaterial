package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Unmarshal("KEY=value")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err2 := godotenv.Write(env, "./.new-env")
	if err2 != nil {
		log.Fatal("Error writing .env file")
	}

	// writing to a string
	env, err = godotenv.Unmarshal("KEY=value")
	if err != nil {
		log.Fatal("Error writing .env string")
	}
	content, err := godotenv.Marshal(env)
	fmt.Println("content:", content)
}

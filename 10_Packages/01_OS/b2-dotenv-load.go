package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// go get -u github.com/joho/godotenv

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	fmt.Printf("%s uses %s\n", os.Getenv("NAME"), os.Getenv("EDITOR"))
}

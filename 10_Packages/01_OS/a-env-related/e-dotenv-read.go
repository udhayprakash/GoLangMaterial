package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// go install github.com/joho/godotenv@latest

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	fmt.Printf("%s uses %s\n", os.Getenv("NAME"), os.Getenv("EDITOR"))
}

// Assignment - research how to use .env file present in one or two levels above in the paths

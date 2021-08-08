package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// go get -u github.com/joho/godotenv
func main() {
	err := godotenv.Load(".env2")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Printf(`
		s3Bucket : %s
		secretKey: %s
	Person Details
		Name     : %s
		Age      : %s
		height   : %s
	`, s3Bucket, secretKey, 
		os.Getenv("NAME"), os.Getenv("AGE"), 
		os.Getenv("HEIGHT"))
	// All loaded content will be of string type

	// To load all .env variables into a map,
	var myEnv map[string]string
	myEnv, err1 := godotenv.Read()
	if err1 != nil {
		log.Fatal("Error reading .env file")
	}
	fmt.Println(myEnv)

}
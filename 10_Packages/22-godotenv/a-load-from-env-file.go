package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// go get -u github.com/joho/godotenv
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Printf(`
		s3Bucket : %s
		secretKey: %s
	`, s3Bucket, secretKey)
}

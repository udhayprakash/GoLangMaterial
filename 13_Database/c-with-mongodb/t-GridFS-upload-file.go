package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	bucket := db.GridFSBucket()

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	uploadStream, err := bucket.OpenUploadStream(context.TODO(), "example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadStream.Close()

	if _, err := uploadStream.Write([]byte("Hello, GridFS!")); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded successfully")
}

// mongofiles -d testdb put example.txt

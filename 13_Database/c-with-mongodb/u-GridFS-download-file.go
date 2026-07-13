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

	file, err := os.Create("downloaded_example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n, err := bucket.DownloadToStreamByName(context.TODO(), "example.txt", file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Downloaded %d bytes successfully\n", n)
}

// mongofiles -d testdb get example.txt

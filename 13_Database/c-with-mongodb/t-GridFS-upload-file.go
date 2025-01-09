package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/gridfs"
    "os"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    db := client.Database("testdb")
    bucket, err := gridfs.NewBucket(db)
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    uploadStream, err := bucket.OpenUploadStream("example.txt")
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

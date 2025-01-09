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

    downloadStream, err := bucket.OpenDownloadStreamByName("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer downloadStream.Close()

    file, err := os.Create("downloaded_example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if _, err := downloadStream.WriteTo(file); err != nil {
        log.Fatal(err)
    }

    fmt.Println("File downloaded successfully")
}

// mongofiles -d testdb get example.txt

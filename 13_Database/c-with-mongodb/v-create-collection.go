package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    db := client.Database("testdb")
    err = db.CreateCollection(context.TODO(), "newcollection")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Collection created successfully")
}

// db.createCollection("newcollection")

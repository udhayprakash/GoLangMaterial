package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("testdb").Collection("users")
    changeStream, err := collection.Watch(context.TODO(), mongo.Pipeline{})
    if err != nil {
        log.Fatal(err)
    }
    defer changeStream.Close(context.TODO())

    for changeStream.Next(context.TODO()) {
        var change bson.M
        if err := changeStream.Decode(&change); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Change detected: %v\n", change)
    }
}

// db.users.watch()

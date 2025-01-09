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

    collection := client.Database("testdb").Collection("users")
    cursor, err := collection.Indexes().List(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.TODO())

    var results []bson.M
    if err = cursor.All(context.TODO(), &results); err != nil {
        log.Fatal(err)
    }

    for _, result := range results {
        fmt.Printf("Found index: %v\n", result)
    }
}
// db.users.getIndexes()

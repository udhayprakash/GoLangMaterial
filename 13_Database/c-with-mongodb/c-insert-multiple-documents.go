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
    docs := []interface{}{
        bson.D{{"name", "Alice"}, {"age", 25}},
        bson.D{{"name", "Bob"}, {"age", 30}},
    }

    insertResult, err := collection.InsertMany(context.TODO(), docs)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted documents with IDs:", insertResult.InsertedIDs)
}

// db.users.insertMany([{name: "Alice", age: 25}, {name: "Bob", age: 30}])

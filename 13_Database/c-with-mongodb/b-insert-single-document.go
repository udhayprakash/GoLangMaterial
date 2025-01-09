package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("testdb").Collection("users")
    doc := bson.D{{"name", "John Doe"}, {"age", 30}}

    insertResult, err := collection.InsertOne(context.TODO(), doc)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted document with ID:", insertResult.InsertedID)
}

// db.users.insertOne({name: "John Doe", age: 30})

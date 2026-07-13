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
    filter := bson.D{{"name", "John Doe"}}

    deleteResult, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Deleted %v documents.\n", deleteResult.DeletedCount)
}

// db.users.deleteOne({name: "John Doe"})

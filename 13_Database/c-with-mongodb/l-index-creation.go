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
    indexModel := mongo.IndexModel{
        Keys: bson.D{{"name", 1}},
    }

    indexName, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created index %v\n", indexName)
}

// db.users.createIndex({name: 1})

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

    databases, err := client.ListDatabaseNames(context.TODO(), bson.D{})
    if err != nil {
        log.Fatal(err)
    }

    for _, db := range databases {
        fmt.Println(db)
    }
}

// show dbs

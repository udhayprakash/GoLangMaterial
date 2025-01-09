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
    _, err = collection.Indexes().DropOne(context.TODO(), "name_1")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Dropped index name_1")
}
// db.users.dropIndex("name_1")

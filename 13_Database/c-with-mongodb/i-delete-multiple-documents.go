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
    filter := bson.D{{"age", bson.D{{"$gt", 30}}}}

    deleteResult, err := collection.DeleteMany(context.TODO(), filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Deleted %v documents.\n", deleteResult.DeletedCount)
}

// db.users.deleteMany({age: {$gt: 30}})

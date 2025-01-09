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
    filter := bson.D{{"age", bson.D{{"$gt", 25}}}}
    update := bson.D{{"$inc", bson.D{{"age", 1}}}}

    updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// db.users.updateMany({age: {$gt: 25}}, {$inc: {age: 1}})

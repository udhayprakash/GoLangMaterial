package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
    filter := bson.D{{"age", bson.D{{"$gt", 25}}}}

    count, err := collection.CountDocuments(context.TODO(), filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %v documents.\n", count)
}
// db.users.countDocuments({age: {$gt: 25}})

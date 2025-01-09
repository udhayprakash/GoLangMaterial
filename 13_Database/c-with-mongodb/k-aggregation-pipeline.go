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
	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"age", bson.D{{"$gt", 25}}}}}},
		bson.D{{"$group", bson.D{{"_id", "$name"}, {"total", bson.D{{"$sum", 1}}}}}},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Printf("Found document: %v\n", result)
	}
}

/*
	db.users.aggregate([
		{$match: {age: {$gt: 25}}},
		{$group: {_id: "$name", total: {$sum: 1}}}
	])

*/

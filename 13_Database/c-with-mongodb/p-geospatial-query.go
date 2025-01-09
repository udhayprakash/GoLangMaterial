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

	collection := client.Database("testdb").Collection("places")
	filter := bson.D{
		{"location", bson.D{
			{"$near", bson.D{
				{"$geometry", bson.D{
					{"type", "Point"},
					{"coordinates", []float64{-73.9667, 40.78}},
				}},
				{"$maxDistance", 1000},
			}},
		}},
	}

	cursor, err := collection.Find(context.TODO(), filter)
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
db.places.find({
    location: {
        $near: {
            $geometry: {
                type: "Point",
                coordinates: [-73.9667, 40.78]
            },
            $maxDistance: 1000
        }
    }
})

*/
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
    models := []mongo.WriteModel{
        mongo.NewInsertOneModel().SetDocument(bson.D{{"name", "Alice"}, {"age", 25}}),
        mongo.NewUpdateOneModel().SetFilter(bson.D{{"name", "Bob"}}).SetUpdate(bson.D{{"$set", bson.D{{"age", 31}}}}),
        mongo.NewDeleteOneModel().SetFilter(bson.D{{"name", "Charlie"}}),
    }

    bulkResult, err := collection.BulkWrite(context.TODO(), models)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Inserted %v, Updated %v, Deleted %v\n", bulkResult.InsertedCount, bulkResult.ModifiedCount, bulkResult.DeletedCount)
}
/*
db.users.bulkWrite([
    {insertOne: {document: {name: "Alice", age: 25}}},
    {updateOne: {filter: {name: "Bob"}, update: {$set: {age: 31}}}},
    {deleteOne: {filter: {name: "Charlie"}}}
])

*/
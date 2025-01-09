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

    db := client.Database("admin")
    command := bson.D{{"usersInfo", 1}}
    var result bson.M
    err = db.RunCommand(context.TODO(), command).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Users: %v\n", result["users"])
}

// db.getUsers()

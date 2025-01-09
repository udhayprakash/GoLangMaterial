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

    session, err := client.StartSession()
    if err != nil {
        log.Fatal(err)
    }
    defer session.EndSession(context.TODO())

    err = mongo.WithSession(context.TODO(), session, func(sessionContext mongo.SessionContext) error {
        if err = session.StartTransaction(); err != nil {
            return err
        }

        collection := client.Database("testdb").Collection("users")
        if _, err = collection.InsertOne(sessionContext, bson.D{{"name", "Alice"}, {"age", 25}}); err != nil {
            session.AbortTransaction(sessionContext)
            return err
        }

        if _, err = collection.InsertOne(sessionContext, bson.D{{"name", "Bob"}, {"age", 30}}); err != nil {
            session.AbortTransaction(sessionContext)
            return err
        }

        return session.CommitTransaction(sessionContext)
    })

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Transaction committed successfully")
}

/*/
	session.startTransaction()
	db.users.insertOne({name: "Alice", age: 25})
	db.users.insertOne({name: "Bob", age: 30})
	session.commitTransaction()
*/
package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    stmt, err := db.Prepare("INSERT INTO users(username, email) VALUES(?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Insert multiple users
    users := []struct {
        username string
        email    string
    }{
        {"user1", "user1@example.com"},
        {"user2", "user2@example.com"},
        {"user3", "user3@example.com"},
    }

    for _, user := range users {
        _, err := stmt.Exec(user.username, user.email)
        if err != nil {
            log.Printf("Error inserting user %s: %v\n", user.username, err)
            continue
        }
        log.Printf("Inserted user: %s\n", user.username)
    }
}
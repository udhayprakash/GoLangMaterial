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

    insertSQL := `
    INSERT INTO users (username, email)
    VALUES (?, ?)`

    result, err := db.Exec(insertSQL, "johndoe", "john@example.com")
    if err != nil {
        log.Fatal(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Inserted user with ID: %d\n", id)
}
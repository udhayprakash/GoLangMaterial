package main

import (
    "context"
    "database/sql"
    "log"
    "time"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var username string
    err = db.QueryRowContext(ctx, "SELECT username FROM users WHERE id = ?", 1).Scan(&username)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Username: %s\n", username)
}
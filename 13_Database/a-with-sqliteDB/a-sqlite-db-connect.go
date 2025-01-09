package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

// GCC in widnows - https://jmeubank.github.io/tdm-gcc/

func main() {
    // Open an in-memory SQLite database
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a table
    _, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
    if err != nil {
        log.Fatal(err)
    }

    // Insert data
    _, err = db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob')")
    if err != nil {
        log.Fatal(err)
    }

    // Query data
    rows, err := db.Query("SELECT id, name FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }
}

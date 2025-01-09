package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID       int
    Username string
    Email    string
}

func main() {
    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var user User
    query := "SELECT id, username, email FROM users WHERE id = ?"
    err = db.QueryRow(query, 1).Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("No user found")
            return
        }
        log.Fatal(err)
    }

    log.Printf("User found: %+v\n", user)
}
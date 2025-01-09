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

    rows, err := db.Query("SELECT id, username, email FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Username, &user.Email)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    for _, user := range users {
        log.Printf("User: %+v\n", user)
    }
}
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

    updateSQL := `
    UPDATE users 
    SET email = ? 
    WHERE id = ?`

    result, err := db.Exec(updateSQL, "newemail@example.com", 1)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Updated %d rows\n", rowsAffected)
}
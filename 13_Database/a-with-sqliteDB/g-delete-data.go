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

    deleteSQL := `DELETE FROM users WHERE id = ?`

    result, err := db.Exec(deleteSQL, 1)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Deleted %d rows\n", rowsAffected)
}

package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./banking.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create accounts table if not exists
    _, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS accounts (
        id INTEGER PRIMARY KEY,
        balance DECIMAL(10,2)
    )`)
    if err != nil {
        log.Fatal(err)
    }

    // Start transaction
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }

    // Transfer $100 from account 1 to account 2
    _, err = tx.Exec("UPDATE accounts SET balance = balance - 100 WHERE id = 1")
    if err != nil {
        tx.Rollback()
        log.Fatal(err)
    }

    _, err = tx.Exec("UPDATE accounts SET balance = balance + 100 WHERE id = 2")
    if err != nil {
        tx.Rollback()
        log.Fatal(err)
    }

    // Commit transaction
    err = tx.Commit()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Transaction completed successfully")
}
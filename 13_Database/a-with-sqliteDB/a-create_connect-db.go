package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")
}

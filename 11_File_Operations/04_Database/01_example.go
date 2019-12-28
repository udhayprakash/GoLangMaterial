package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/DjProject1")
	//"user:password@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
// Ref: http://go-database-sql.org/accessing.html
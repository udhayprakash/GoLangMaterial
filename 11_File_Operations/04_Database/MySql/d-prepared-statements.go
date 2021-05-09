package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
With prepared statements, we use placeholders instead of
directly writing the values into the statements.

Prepared statements increase security and performance of
database operations.
*/
type City struct {
	Id         int
	Name       string
	Population int
}

func main() {

	db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var myid int = 1

	res, err := db.Query("SELECT * FROM cities WHERE id = ?", myid)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {

		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", city)
	} else {

		fmt.Println("No city found")
	}
}

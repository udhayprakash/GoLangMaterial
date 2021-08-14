package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sample-fb/config"
)

var (
	db  *sql.DB
	err error
)

//NewClient is used to establish the connection between postgresql and our services
func NewClient() {
	//enable database connection
	db, err = sql.Open(config.DBName, fmt.Sprintf("host=% port=% user=% password=% dbname=% sslmode=%", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully established DB Connection")
}

package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
	go get gorm.io/driver/sqlite
    go get gorm.io/gorm
*/
type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
}

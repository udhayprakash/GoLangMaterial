package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

	tx := db.Begin()
	if err := tx.Create(&User{Name: "David", Age: 40}).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

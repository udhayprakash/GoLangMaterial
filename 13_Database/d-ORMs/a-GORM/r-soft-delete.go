package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Age       int
	DeletedAt gorm.DeletedAt
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	var user User
	db.First(&user, 1)
	db.Delete(&user)
}

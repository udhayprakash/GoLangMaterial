package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CreditCard struct {
	ID     uint
	Number string
	UserID uint
}

type User struct {
	ID         uint
	Name       string
	Age        int
	CreditCard []CreditCard
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &CreditCard{})

	user := User{Name: "Frank", Age: 45, CreditCard: []CreditCard{{Number: "1234-5678-9012-3456"}}}
	db.Create(&user)
}

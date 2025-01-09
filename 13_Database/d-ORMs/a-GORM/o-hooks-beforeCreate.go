package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before creating user:", u.Name)
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&User{Name: "Eve", Age: 28})
}

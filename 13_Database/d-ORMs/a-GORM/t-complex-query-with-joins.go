package main


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

    var users []User
    db.Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "1234-5678-9012-3456").Find(&users)
    fmt.Println(users)
}

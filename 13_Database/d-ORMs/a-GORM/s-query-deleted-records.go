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

    var user User
    db.Unscoped().Where("name = ?", "John").Find(&user)
    fmt.Println(user)
}
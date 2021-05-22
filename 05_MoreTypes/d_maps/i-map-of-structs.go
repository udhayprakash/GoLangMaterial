package main

import "fmt"

type User struct {
	name       string
	occupation string
}

func main() {

	u1 := User{
		name:       "John Doe",
		occupation: "gardener",
	}

	u2 := User{
		name:       "Richard Roe",
		occupation: "driver",
	}

	users := map[int]User{
		1: u1,
		2: u2,
		3: User{
			name:       "Lucy Smith",
			occupation: "teacher",
		},
	}

	fmt.Println(users)
}

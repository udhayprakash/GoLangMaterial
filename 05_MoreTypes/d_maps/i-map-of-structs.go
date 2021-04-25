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

	u3 := User{
		name:       "Lucy Smith",
		occupation: "teacher",
	}

	users := map[int]User{
		1: u1,
		2: u2,
		3: u3,
	}

	fmt.Println(users)
}

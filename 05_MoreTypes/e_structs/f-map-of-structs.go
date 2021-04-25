package main

import "fmt"

type User struct {
	name       string
	occupation string
}

type Thing struct {
	Foo string
	Bar int
	Baz float32
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
	}

	fmt.Println(users)

	mapOfStructs := map[string]Thing{
		"default values": Thing{},
		"positional":     Thing{"udhay", 123, 11.123},
		"keyBased":       Thing{Foo: "udhay", Bar: 123, Baz: 11.123},
	}
	fmt.Println("mapOfStructs=", mapOfStructs)
}

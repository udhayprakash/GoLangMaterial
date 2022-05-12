package main

import "fmt"

func main() {

	type User struct {
		name       string
		occupation string
	}

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
			name:       "Modi",
			occupation: "Prime minister",
		},
		4: {
			name:       "kovid",
			occupation: "president",
		},
	}
	fmt.Println(users)
	// map[1:{John Doe gardener} 2:{Richard Roe driver} 3:{Modi Prime minister} 4:{kovid president}]
	fmt.Printf("%#v\n", users)
	// map[int]main.User{1:main.User{name:"John Doe", occupation:"gardener"},
	//                   2:main.User{name:"Richard Roe", occupation:"driver"},
	//                   3:main.User{name:"Modi", occupation:"Prime minister"},
	//                   4:main.User{name:"kovid", occupation:"president"}}

}

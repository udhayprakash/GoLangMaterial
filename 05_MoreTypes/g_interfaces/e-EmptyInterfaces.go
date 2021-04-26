package main

import (
	"fmt"
)

type animal struct {
	age   int
	color string
	sound string
}

type cat struct {
	animal
	bossy bool
}

type dog struct {
	animal
	friendly bool
}

func result(a interface{}) { // Here!
	fmt.Println(a)
}

func main() {
	cassy := cat{animal{3, "white", "meow"}, true}
	kitty := cat{animal{2, "black", "meow"}, true}
	missy := cat{animal{1, "brown", "meow"}, true}

	rocky := dog{animal{4, "brown", "woof"}, true}
	lucy := dog{animal{3, "white", "woof"}, true}
	tucker := dog{animal{2, "black", "woof"}, true}

	result(cassy)
	result(rocky)

	// Empty Interface as slice
	friends := []interface{}{cassy, missy, kitty, rocky, lucy, tucker}
	fmt.Println(friends)
}

// {{3 white meow} true}
// {{4 brown woof} true}

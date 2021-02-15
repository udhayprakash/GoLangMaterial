package main

import (
	"container/vector"
)

func main() {
	k1 := vector.IntVector{}
	k2 := &vector.IntVector{}
	k3 := new(vector.IntVector)
	k1.Push(2)
	k2.Push(3)
	k3.Push(4)
}

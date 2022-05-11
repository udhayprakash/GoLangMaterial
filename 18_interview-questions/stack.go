package main

import (
	"fmt"
)

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func main() {
	// Integers stack
	myStackOfInts := new(StackOfInts)
	fmt.Println("myStackOfInts:", myStackOfInts)

	myStackOfInts.Push(123)
	myStackOfInts.Push(456)
	myStackOfInts.Push(789)

	value, _ := myStackOfInts.Pop()
	fmt.Println(value)
	value, _ = myStackOfInts.Pop()
	fmt.Println("myStackOfInts.IsEmpty():", myStackOfInts.IsEmpty())
	fmt.Println("myStackOfInts:", myStackOfInts)

	// strings stack
	myStackOfStrings := new(StackOfStrings)
	fmt.Println("myStackOfStrings.IsEmpty():", myStackOfStrings.IsEmpty())

	myStackOfStrings.Push("one two three")
	myStackOfStrings.Push("four five six")

	strValue, _ := myStackOfStrings.Pop()
	fmt.Println(strValue)

}

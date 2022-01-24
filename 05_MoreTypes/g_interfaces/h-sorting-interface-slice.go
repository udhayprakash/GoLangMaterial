package main

import (
	"fmt"
	"sort"
)

type GoLangStudent struct {
	Name  string
	Marks int
}

func (s GoLangStudent) String() string {
	return fmt.Sprintf("%s: %d", s.Name, s.Marks)
}

type ByMarks []GoLangStudent // slice of structs

// Len, Swap and Less methods , with same name, need to be defined, for sorting
func (b ByMarks) Len() int      { return len(b) }
func (b ByMarks) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByMarks) Less(i, j int) bool {
	return b[i].Marks < b[j].Marks
}

func main() {
	students := []GoLangStudent{
		{"appanna", 31},
		{"Jay", 42},
		{"Mani", 17},
		{"Jesudas", 26},
	}
	fmt.Println("Before sorting =", students)
	// Method1
	sort.Sort(ByMarks(students))
	fmt.Println("After sorting  =", students)

	sort.Sort(sort.Reverse(ByMarks(students)))
	fmt.Println("After reversing=", students)

	// Method2
	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks < students[j].Marks
	})
	fmt.Println("After sorting  =", students)
}

// assignment - try enhancing this custom sort, to sort by length of student name

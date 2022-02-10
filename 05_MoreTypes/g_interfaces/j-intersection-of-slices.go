package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6}

	fmt.Println("Intersection(slice1, slice2)        = ", Intersection(slice1, slice2))
	fmt.Println("IntersectionAnyType(slice1, slice2) = ", IntersectionAnyType(slice1, slice2))

}

func Intersection(s1, s2 []int) []int {
	s1Map := make(map[int]int)
	for _, num1 := range s1 {
		s1Map[num1] = 1
	}
	var intersection = []int{}
	for _, num2 := range s2 {
		_, isPresent := s1Map[num2]
		if isPresent == true {
			intersection = append(intersection, num2)
		}
	}
	return intersection
}

func IntersectionAnyType(s1, s2 interface{}) interface{} {
	s1Map := make(map[interface{}]int)

	switch reflect.TypeOf(s1).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(s1)

		for i := 0; i < s.Len(); i++ {
			num1 := s.Index(i).Int()
			s1Map[num1] = 1
		}
	}
	// fmt.Println("s1Map =", s1Map)

	var intersection = []int64{}

	switch reflect.TypeOf(s2).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(s2)

		for i := 0; i < s.Len(); i++ {
			num2 := s.Index(i).Int()
			_, isPresent := s1Map[num2]
			if isPresent == true {
				intersection = append(intersection, num2)
			}
		}
	}

	return intersection
}

package main

import (
	"fmt"
	"reflect"
)

/*
Problem: Remove duplicate values from a given sorted
		singly linked list consisting of N nodes.

Input:
	LinkedList	: 0->4->4->5
	Output		: 0 4 5
*/

type Node struct {
	value int
	next  *Node
}

type SingleLinkedList struct {
	head   *Node
	length int
}

// methods

func (sl *SingleLinkedList) Insert(val int) {
	n1 := Node{value: val}

	if sl.length == 0 {
		sl.head = &n1
		sl.length++
	}

	ptr := sl.head
	for i := 0; i < sl.length; i++ {
		if ptr.next == nil {
			ptr.next = &n1
			sl.length++
			return
		}
		ptr = ptr.next
	}
}

func (sl *SingleLinkedList) GetUniqueValues() {
	if sl.length == 0 {
		fmt.Println("No nodes in the linked list")
		return
	}
	ptr := sl.head
	unique := make(map[int]int)
	for i := 0; i < sl.length; i++ {
		fmt.Println("ptr", ptr, " *ptr", *ptr, "ptr.value", ptr.value)
		unique[ptr.value]++ // freq
		ptr = ptr.next
	}
	fmt.Println(reflect.ValueOf(unique).MapKeys()) //
}

func (sl *SingleLinkedList) DeleteDups() {
	if sl.length == 0 {
		fmt.Println("No nodes in the linked list")
		return
	}
	// [0  1 2 2 3 4 5 6 6 7 8]
	ptr := sl.head // 0
	var prev, curr int
	for i := 0; i < sl.length; i++ {
		if i == 0 {
			prev, curr = ptr.value, ptr.value // prev=0, curr = 0
		} else {
			curr = ptr.value // prev =2, curr = 2

			if prev == curr {
				prev = ptr.value //    prev =2, curr =2
				// delet a node
				ptr.next = ptr.next.ptr.next
			} else {
				prev = ptr.value //    prev =2, curr =2
			}
		}

		ptr = ptr.next
	}
}

func main() {
	myslice := []int{0, 1, 2, 2, 3, 4, 4, 4, 5}

	sll := SingleLinkedList{}

	for _, vl := range myslice {
		sll.Insert(vl)
	}

	// Get Unique values
	sll.GetUniqueValues()
}

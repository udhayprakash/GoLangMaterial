package main

import (
	"fmt"
)

type node struct {
	value int
	prev  *node
	next  *node
}

var head *node = nil

func (list *node) pushBack(value int) *node {
	if head == nil {
		list.value = value
		list.next = nil
		list.prev = nil
		head = list
		return list
	} else {
		for list.next != nil {
			if list.value == value {
				return list
			}
			list = list.next
		}
		list.next = new(node)
		list.next.value = value
		list.next.prev = list
		list.next.next = nil
		return list
	}
}

func (list *node) pushFront(value int) *node {
	if head == nil {
		list.value = value
		list.next = nil
		list.prev = nil
		head = list
		return list
	} else {
		tmp := new(node)
		tmp.next = list
		tmp.prev = nil
		tmp.value = value
		head = tmp
		return tmp
	}
}

func (list *node) popFront() *node {
	if head != nil {
		tmp := new(node)
		tmp = head.next
		tmp.prev = nil
		head = tmp
		return list
	}
	return list
}

func (list *node) popBack() *node {
	if head == nil {
		return list
	}
	tmp := new(node)
	tmp = list
	for tmp.next.next != nil {
		tmp = tmp.next
	}
	tmp.next = nil
	return list
}

func (list *node) printElements() {
	for head.next != nil {
		fmt.Println("Yes")
		fmt.Println(head.value)
		head = head.next
	}
	fmt.Println(head.value)
}

func main() {
	listval := new(node)
	listval.pushBack(22)
	listval.pushBack(12)
	listval.pushFront(1)
	listval.popFront()
	listval.popBack()
	dup := head
	for dup != nil {
		fmt.Printf(" Current element val : %d ", dup.value)
		fmt.Printf("Prev element val : %d ", dup.prev)
		fmt.Printf("Next element val : %d ", dup.next)
		dup = dup.next
	}
	fmt.Println("Value :", head)
}

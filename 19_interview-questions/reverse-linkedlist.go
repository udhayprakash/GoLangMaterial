package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v ", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}
func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 0 {
		return
	}
	link := List{}
	s := make([]int, n)
	for i := range s {
		fmt.Scanln(&s[i])
		link.Insert(s[i])
	}
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Printf("head: %v\n", link.tail.key)
	fmt.Printf("tail: %v\n", link.head.key)
	link.Reverse()
}

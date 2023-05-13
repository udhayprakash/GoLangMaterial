package main

import "fmt"

type node struct {
	val  int
	next *node
	prev *node
}

type circularlinkedlist struct {
	head *node
	tail *node
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val int) *node {
	n := &node{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *circularlinkedlist) addAtBeg(val int) {
	n := newNode(val)
	n.next = ll.head
	ll.head = n
	if ll.tail == nil {
		ll.tail = n
	}
	n.prev = ll.tail
	ll.tail.next = n
}

func (ll *circularlinkedlist) addAtEnd(val int) {
	n := newNode(val)
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return
	}
	if ll.head == ll.tail {
		ll.head.next = n
		ll.head.prev = n
		n.next = ll.head
		n.prev = ll.head
		ll.tail = n
		return
	}
	cur := ll.head
	for ; cur.next != ll.tail; cur = cur.next {
	}
	cur.next.next = n
	n.prev = cur.next
	n.next = ll.head
	ll.tail = n
}

func (ll *circularlinkedlist) delAtBeg() int {
	if ll.head == nil {
		return -1
	}
	if ll.head == ll.tail {
		value := ll.head
		ll.head = nil
		ll.tail = nil
		return value.val
	}
	cur := ll.head
	ll.head = cur.next
	if ll.head != nil {
		ll.head.prev = ll.tail
		ll.tail.next = ll.head
	}
	return cur.val
}

func (ll *circularlinkedlist) delAtEnd() int {
	// no item
	if ll.head == nil {
		return -1
	}
	// only one item
	if ll.head == ll.tail {
		return ll.delAtBeg()
	}
	// more than one, go to second last
	cur := ll.head
	for ; cur.next != ll.tail; cur = cur.next {
	}
	retval := cur.next.val
	cur.next = ll.head
	ll.tail = cur
	ll.head.prev = cur
	return retval
}

func (ll *circularlinkedlist) count() int {
	var ctr int = 0
	for cur := ll.head; cur != ll.tail; cur = cur.next {
		ctr++
	}
	return ctr
}

func (ll *circularlinkedlist) reverse() {
	var prev, next *node
	if ll.head == ll.tail {
		return
	}
	cur := ll.head.next
	ll.tail = ll.head
	if ll.head.next != nil {
		for cur != ll.head {
			next = cur.next
			cur.next = prev
			cur.prev = next
			prev = cur
			cur = next
		}
	}
	ll.head = prev
	cur.next = ll.tail.next
	ll.tail.next.next = ll.tail
	ll.tail.next = ll.head
	ll.head.prev = ll.tail
}

func (ll *circularlinkedlist) display() {
	if ll.head != nil {
		fmt.Print(ll.head.val, " ")
		if ll.head.next != nil {
			for cur := ll.head.next; cur != ll.head; cur = cur.next {
				fmt.Print(cur.val, " ")
			}
		}
	}
	fmt.Print("")
}

func (ll *circularlinkedlist) displayReverse() {
	if ll.head == nil {
		return
	}
	var cur *node
	for cur = ll.head; cur.next != ll.tail; cur = cur.next {
	}
	for ; cur != ll.tail; cur = cur.prev {
		fmt.Print(cur.val, " ")
	}
	fmt.Print("")
}

func main() {
	ll := circularlinkedlist{}
	ll.addAtBeg(10)
	ll.addAtEnd(20)
	ll.reverse()
	ll.display()
	fmt.Print("Deleting an element at the beginning : ", ll.delAtBeg(), "")
	fmt.Print("Deleting an element at the end : ", ll.delAtEnd(), "")
	ll.display()
}

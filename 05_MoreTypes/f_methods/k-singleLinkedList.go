package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func CreateNode(value int) *Node {
	var n Node
	n.value = value
	n.next = nil
	return &n
}

func TraverseLinkedList(head *Node) {
	fmt.Print("Linked list: ")
	for temp := head; temp != nil; {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
}

func main() {
	head := CreateNode(0)
	head.next = CreateNode(10)
	head.next.next = CreateNode(20)
	head.next.next.next = CreateNode(30)
	TraverseLinkedList(head)

}

/*

(0, addrss ) -> (10, address)  -> (20, address)  -> (30, address)

*/

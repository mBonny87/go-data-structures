package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) InsertAtBeginning(v int) {
	if dll.head == nil {
		dll.head = &Node{value: v, next: nil, prev: nil}
		dll.tail = &Node{value: v, next: nil, prev: nil}
		return
	}
	temp := dll.head
	new := &Node{value: v, next: temp, prev: nil}
	dll.head = new

	for temp.next != nil {
		temp = temp.next
	}
	dll.tail = temp
}

func (dll *DoublyLinkedList) Print() {
	curr := dll.head
	for curr != nil {
		fmt.Println("Value:", curr.value)
		curr = curr.next
	}
}

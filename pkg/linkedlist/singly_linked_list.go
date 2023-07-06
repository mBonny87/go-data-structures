package linkedlist

import "fmt"

type SinglyLinkedList struct {
	head *Node
}

func (ll *SinglyLinkedList) Insert(v int) {
	new := &Node{Value: v, Next: nil}

	if ll.head == nil {
		ll.head = new
		return
	}

	curr := ll.head //pointer
	for curr.Next != nil {
		curr = curr.Next // Next pointer
	}

	curr.Next = new
}

func (ll *SinglyLinkedList) Search(v int) bool {
	curr := ll.head
	if curr == nil {
		return false
	}
	for curr != nil {
		if v == curr.Value {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (ll *SinglyLinkedList) Delete(v int) {
	if ll.head == nil {
		fmt.Println("Linked List is empty, cannot delete this Value")
		return
	}
	if ll.head.Value == v {
		ll.head = ll.head.Next
		return
	}

	curr := ll.head
	for curr.Next != nil {
		if curr.Next.Value == v {
			curr.Next = curr.Next.Next
			return
		}

		curr = curr.Next
	}

}

func (ll *SinglyLinkedList) Length() int {
	count := 0
	curr := ll.head

	for curr != nil {
		count++
		curr = curr.Next
	}

	return count

}

func (ll *SinglyLinkedList) Reverse() SinglyLinkedList {
	if ll.head == nil {
		fmt.Println("Cannot reverse empty linked list")
		return *ll
	}
	if ll.head.Next == nil {
		return *ll
	}

	var prev *Node
	curr := ll.head
	var Next *Node

	for curr != nil {
		Next = curr.Next
		curr.Next = prev
		prev = curr
		curr = Next
	}

	ll.head = prev

	return *ll
}

func (ll *SinglyLinkedList) Print() {
	curr := ll.head
	for curr != nil {
		fmt.Println("Value:", curr.Value)
		curr = curr.Next
	}
}

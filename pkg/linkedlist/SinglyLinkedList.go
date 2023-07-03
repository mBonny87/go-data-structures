package linkedlist

import "fmt"

type SinglyLinkedList struct {
	head *Node
}

func (ll *SinglyLinkedList) Insert(v int) {
	new := &Node{value: v, next: nil}

	if ll.head == nil {
		ll.head = new
		return
	}

	curr := ll.head //pointer
	for curr.next != nil {
		curr = curr.next // next pointer
	}

	curr.next = new
}

func (ll *SinglyLinkedList) Print() {
	curr := ll.head
	if curr == nil {
		fmt.Println("Empty SinglyLinkedList")
		return
	}
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (ll *SinglyLinkedList) Search(v int) bool {
	curr := ll.head
	if curr == nil {
		return false
	}
	for curr != nil {
		if v == curr.value {
			return true
		}
		curr = curr.next
	}
	return false
}

func (ll *SinglyLinkedList) Delete(v int) {
	if ll.head == nil {
		fmt.Println("Linked List is empty, cannot delete this value")
		return
	}
	if ll.head.value == v {
		ll.head = ll.head.next
		return
	}

	curr := ll.head
	for curr.next != nil {
		if curr.next.value == v {
			curr.next = curr.next.next
			return
		}

		curr = curr.next
	}

}

func (ll *SinglyLinkedList) Length() int {
	count := 0
	curr := ll.head

	for curr != nil {
		count++
		curr = curr.next
	}

	return count

}

func (ll *SinglyLinkedList) Reverse() SinglyLinkedList {
	if ll.head == nil {
		fmt.Println("Cannot reverse empty linked list")
		return *ll
	}
	if ll.head.next == nil {
		return *ll
	}

	var prev *Node
	curr := ll.head
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	ll.head = prev

	return *ll
}

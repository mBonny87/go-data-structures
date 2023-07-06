package linkedlist

import (
	"fmt"
)

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) Shift(v int) {
	if dll.head == nil {
		dll.head = &Node{Value: v, Prev: nil, Next: nil}
		return
	}
	if dll.head.Next == nil {
		dll.tail = dll.head
		dll.head = &Node{Value: v, Prev: nil, Next: dll.tail}
		dll.tail.Prev = dll.head
		return
	}

	new := &Node{Value: v, Prev: nil, Next: dll.head}
	dll.head.Prev = new
	dll.head = new
}

func (dll *DoublyLinkedList) Push(v int) {
	if dll.head == nil {
		dll.head = &Node{Value: v, Prev: nil, Next: nil}
		return
	}
	if dll.head.Next == nil {
		dll.tail = &Node{Value: v, Prev: dll.head, Next: nil}
		dll.head.Next = dll.tail
		return
	}

	new := &Node{Value: v, Prev: dll.tail, Next: nil}
	dll.tail.Next = new
	dll.tail = new

}

func (dll *DoublyLinkedList) Unshift() {
	if dll.head == nil {
		return
	}
	if dll.head.Next == nil {
		dll.head = nil
	}

	newHead := dll.head.Next
	newHead.Prev = nil
	dll.head = newHead
}

func (dll *DoublyLinkedList) Pop() {
	if dll.head == nil { // 0 elements
		return
	}
	if dll.head.Next == nil { // 1 element
		dll.head = nil
	}

	//2+ elements
	newTail := dll.tail.Prev
	newTail.Next = nil
	dll.tail = newTail
}

func (dll *DoublyLinkedList) PushAfter(node *Node, v int) {
	if dll.head == nil {
		return
	}

	curr := dll.head
	for curr != nil {
		// if I don't have a next
		if v == curr.Value && curr.Next == nil {
			curr.Next = node
			node.Prev = curr
			dll.tail = node
			break
		}
		if v == curr.Value {
			node.Next = curr.Next
			node.Prev = curr
			curr.Next.Prev = node
			curr.Next = node
			break
		}
		curr = curr.Next
	}

}

func (dll *DoublyLinkedList) Print() {
	curr := dll.head
	for curr != nil {
		fmt.Println("Value:", curr.Value)
		curr = curr.Next
	}
}

func (dll *DoublyLinkedList) DeleteNode(v int) {
	if dll.head == nil {
		return
	}
	if dll.head.Next == nil && dll.head.Value == v {
		dll.head = nil
		dll.tail = nil
		return
	}

	curr := dll.head
	for curr != nil {
		if curr.Value == v {
			if curr.Prev == nil {
				dll.head = dll.head.Next
				dll.head.Prev = nil
				break
			}
			if curr.Next == nil {
				dll.tail = dll.tail.Prev
				dll.tail.Next = nil
				break
			}
			// curr.Prev = curr.Next
			// curr.Next = curr.Prev
			curr.Prev, curr.Next = curr.Next, curr.Prev
			break
		}
		curr = curr.Next
	}
}

func (dll *DoublyLinkedList) ToArray() []int {
	if dll.head == nil {
		return []int{}
	}

	curr := dll.head
	result := []int{}
	for curr != nil {
		result = append(result, curr.Value)

		curr = curr.Next
	}
	return result
}

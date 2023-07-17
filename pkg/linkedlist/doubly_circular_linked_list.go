package linkedlist

import "fmt"

type DoublyCircularLinkedList struct {
	Head *Node
	Tail *Node
}

func (dcll *DoublyCircularLinkedList) Push(v int) {
	if dcll.Head == nil {
		dcll.Head = &Node{Value: v, Prev: nil, Next: nil}
		dcll.Head.Next = dcll.Head
		dcll.Head.Prev = dcll.Head
		dcll.Tail = dcll.Head
		return
	}
	if dcll.Head.Next == dcll.Head {
		newNode := &Node{Value: v, Next: dcll.Head, Prev: dcll.Head}
		dcll.Head.Next = newNode
		dcll.Head.Prev = newNode
		dcll.Tail = newNode
		return
	}
	curr := dcll.Head
	for curr.Next != dcll.Head {
		curr = curr.Next
	}
	newNode := &Node{Value: v, Prev: curr, Next: dcll.Head}
	curr.Next = newNode
	dcll.Tail = newNode
	dcll.Head.Prev = dcll.Tail

}

func (dcll *DoublyCircularLinkedList) Pop() {
	if dcll.Head == nil {
		// empty list
		return
	}
	if dcll.Head.Next == dcll.Head {
		// one element
		dcll.Head = nil
		dcll.Tail = nil
		return
	}
	curr := dcll.Head
	for curr.Next.Next != dcll.Head {
		curr = curr.Next
	}
	curr.Next = dcll.Head
	dcll.Tail = curr
}
func (dcll *DoublyCircularLinkedList) Unshift(v int) {
	if dcll.Head == nil {
		dcll.Head = &Node{Value: v, Prev: nil, Next: nil}
		dcll.Head.Next = dcll.Head
		dcll.Head.Prev = dcll.Head
		dcll.Tail = dcll.Head
		return
	}
	if dcll.Head.Next == dcll.Head {
		dcll.Tail = dcll.Head
		newNode := &Node{Value: v, Next: dcll.Tail, Prev: dcll.Tail}
		dcll.Head = newNode
		return
	}
	newNode := &Node{Value: v, Prev: dcll.Tail, Next: dcll.Head}
	dcll.Head = newNode
	dcll.Tail.Next = dcll.Head
}
func (dcll *DoublyCircularLinkedList) Shift() {
	if dcll.Head == nil {
		// empty list
		return
	}
	if dcll.Head.Next == dcll.Head {
		// one element
		dcll.Head = nil
		dcll.Tail = nil
		return
	}
	dcll.Head = dcll.Head.Next
	dcll.Head.Prev = dcll.Tail
	dcll.Tail.Next = dcll.Head
}

func (dcll DoublyCircularLinkedList) Reverse() {

}

func (dcll *DoublyCircularLinkedList) Print() {
	if dcll.Head == nil {
		return
	}
	if dcll.Head.Next == nil {
		fmt.Println(dcll.Head.Value)
		return
	}
	curr := dcll.Head
	for curr.Next != dcll.Head {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	// print last value
	fmt.Println(curr.Value)
}

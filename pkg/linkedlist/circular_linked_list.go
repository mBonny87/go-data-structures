package linkedlist

import "fmt"

type CircularLinkedList struct {
	Head *Node
}

func (cll *CircularLinkedList) Push(v int) {
	if cll.Head == nil {
		cll.Head = &Node{
			Value: v,
		}
		cll.Head.Next = cll.Head
		return
	}
	curr := cll.Head
	for curr.Next != cll.Head {
		curr = curr.Next
	}
	curr.Next = &Node{
		Value: v,
		Next:  cll.Head,
	}
}

func (cll *CircularLinkedList) Unshift(v int) {
	if cll.Head == nil {
		cll.Head = &Node{
			Value: v,
		}
		cll.Head.Next = cll.Head
		return
	}
	curr := cll.Head
	for curr.Next != cll.Head {
		curr = curr.Next
	}
	curr.Next = &Node{
		Value: v,
		Next:  cll.Head,
	}
	cll.Head = curr.Next
}

func (cll *CircularLinkedList) Pop() bool {
	if cll.Head == nil {
		return false
	}
	if cll.Head.Next == cll.Head {
		cll.Head = nil
		return false
	}

	curr := cll.Head
	for curr.Next.Next != cll.Head {
		curr = curr.Next
	}
	curr.Next = cll.Head
	return true
}

func (cll *CircularLinkedList) Shift() bool {
	if cll.Head == nil {
		return false
	}
	if cll.Head.Next == cll.Head {
		cll.Head = nil
		return false
	}

	oldHead := cll.Head
	cll.Head = cll.Head.Next
	curr := cll.Head
	for curr.Next != oldHead {
		curr = curr.Next
	}
	curr.Next = cll.Head

	return true
}

func (cll *CircularLinkedList) Reverse() bool {
	if cll.Head == nil {
		return false
	}
	if cll.Head.Next == cll.Head {
		return false
	}

	prev := cll.Head

	for prev.Next != cll.Head {
		prev = prev.Next
	}

	curr := cll.Head
	for curr.Next != cll.Head {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr.Next = prev
	cll.Head = curr
	return true
}

func (cll *CircularLinkedList) Print() {
	curr := cll.Head
	for curr.Next != cll.Head {
		fmt.Println("Value:", curr.Value)
		curr = curr.Next
	}
	// print last value
	fmt.Println("Value:", curr.Value)
}

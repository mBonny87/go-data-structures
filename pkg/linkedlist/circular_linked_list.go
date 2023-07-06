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

func (cll *CircularLinkedList) Shift(v int) {
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

func (cll *CircularLinkedList) Print() {
	curr := cll.Head
	for curr.Next != cll.Head {
		fmt.Println("Value:", curr.Value)
		curr = curr.Next
	}
}

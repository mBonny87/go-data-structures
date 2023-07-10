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

// Exercise 2: Implement a method Search(value int) *Node that searches for the first occurrence of the given value in the circular linked list and returns the corresponding node. If the value is not found, the method should return nil.

// Exercise 3: Implement a method Size() int that returns the number of nodes in the circular linked list.

// Exercise 4: Implement a method ToArray() []int that converts the circular linked list to an array and returns it. The array should contain the values of the nodes in the same order as they appear in the circular linked list.

// Exercise 5: Implement a method Reverse() that reverses the order of the nodes in the circular linked list.

func (cll *CircularLinkedList) Print() {
	curr := cll.Head
	for curr.Next != cll.Head {
		fmt.Println("Value:", curr.Value)
		curr = curr.Next
	}
}

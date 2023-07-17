package linkedlist

import (
	"fmt"
	"testing"
)

func TestCircularLinkedList(t *testing.T) {
	cll := CircularLinkedList{}

	// Test push
	cll.Push(10)
	cll.Push(15)
	cll.Push(20)
	cll.Push(25)

	// Check the head
	if cll.Head.Value != 10 {
		t.Errorf("Expected 10 as the head but got %v", cll.Head.Value)
	}

	curr := cll.Head
	for curr.Next != cll.Head {
		curr = curr.Next
	}

	// Check the tail
	if curr.Value != 25 {
		t.Errorf("Expected 25 as the tail but got %v", curr.Value)
	}

	// Chech head after unshift
	cll.Unshift(5)
	if cll.Head.Value != 5 {
		t.Errorf("Expected 5 as the head but got %v", cll.Head.Value)
	}

	// check tail after pop
	if cll.Pop() {
		curr = cll.Head
		for curr != nil && curr.Next != cll.Head {
			curr = curr.Next
		}
		if curr.Value != 20 {
			t.Errorf("Expected 20 as the tail but got %v", curr.Value)
		}
	}

	if cll.Shift() {
		if cll.Head.Value != 10 {
			t.Errorf("Expected 10 as the tail but got %v", cll.Head.Value)
		}
	}

	cll.Print()
	cll.Reverse()
	fmt.Println("Reversed")
	cll.Print()

}

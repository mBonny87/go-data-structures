package linkedlist

import "testing"

func TestCircularLinkedList(t *testing.T) {
	cll := CircularLinkedList{}

	// Test push
	cll.Push(10)
	cll.Push(15)
	cll.Push(20)
	cll.Push(25)

	if cll.Head.Value != 10 {
		t.Errorf("Expected 10 as the head but got %v", cll.Head.Value)
	}

	curr := cll.Head
	for curr.Next != cll.Head {
		curr = curr.Next
	}

	if curr.Value != 25 {
		t.Errorf("Expected 25 as the tail but got %v", curr.Value)
	}

	cll.Shift(5)
	if cll.Head.Value != 5 {
		t.Errorf("Expected 5 as the head but got %v", cll.Head.Value)
	}
}

package linkedlist

import "testing"

func TestSynglyLinkedList(t *testing.T) {
	ll := SinglyLinkedList{}

	// Test Insert
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)
	if ll.head.Value != 5 {
		t.Errorf("Expected 5 as Head but got %v:", ll.head.Value)
	}

	// Test Search
	found := ll.Search(10)
	if !found {
		t.Errorf("Expected 10 to be found")
	}

	// Test Delete
	ll.Delete(15)
	curr := ll.head
	for curr.Next != nil {
		curr = curr.Next
	}
	if curr.Value != 10 {
		t.Errorf("Expected 10 as the Tail but got %v:", curr.Value)
	}

	// Test Length
	if ll.Length() != 2 {
		t.Errorf("Expected length 2 but got %v:", ll.Length())
	}

	ll.Reverse()
	if ll.head.Value != 10 {
		t.Errorf("Expected 10 as Head but got %v:", ll.head.Value)
	}
}

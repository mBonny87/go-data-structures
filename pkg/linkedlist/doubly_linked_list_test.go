package linkedlist

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := DoublyLinkedList{}

	// Test Shift()
	dll.Unshift(10)
	if dll.head.Value != 10 {
		t.Errorf("Expected 10 as Head but got %v:", dll.head.Value)
	}

	// Test Push()
	dll.Push(20)
	dll.Push(30)
	if dll.tail.Value != 30 {
		t.Errorf("Expected 30 as Tail but got %v:", dll.tail.Value)
	}

	// Test Unshift()
	dll.Shift()
	if dll.head.Value != 20 {
		t.Errorf("Expected 20 as Head but got %v:", dll.head.Value)
	}

	// Test Pop()
	dll.Pop()
	if dll.head.Value != 20 || dll.tail.Value != 20 {
		t.Errorf("Expected 30 both Head and Tail but got %v, %v:", dll.head.Value, dll.tail.Value)
	}

	// Test PushAfter()
	dll.PushAfter(&Node{
		Value: 30,
		Prev:  nil,
		Next:  nil,
	}, 20)

	if dll.head.Value != 20 {
		t.Errorf("Expected 20 as Head but got %v:", dll.head.Value)
	}

	dll.DeleteNode(20)

	if dll.head.Next != nil {
		t.Errorf("Expected one element but dll.head.Next but exist")
	}

	mySlice := dll.ToArray()
	if len(mySlice) != 1 {
		t.Errorf("Expected length 1 but go %v", len(mySlice))
	}
}

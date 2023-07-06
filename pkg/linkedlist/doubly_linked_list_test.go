package linkedlist

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := DoublyLinkedList{}

	// Test Shift()
	dll.Shift(10)
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
	dll.Unshift()
	if dll.head.Value != 20 {
		t.Errorf("Expected 20 as Head but got %v:", dll.head.Value)
	}

	// Test Pop()
	dll.Pop()
	if dll.head.Value != 20 || dll.tail.Value != 20 {
		t.Errorf("Expected 30 both Head and Tail but got %v, %v:", dll.head.Value, dll.tail.Value)
	}

	dll.PushAfter(&Node{
		Value: 30,
		Prev:  nil,
		Next:  nil,
	}, 20)

	if dll.head.Value != 20 {
		t.Errorf("Expected 20 as Head but got %v:", dll.head.Value)
	}

	// Test Search()
	// dll.InsertAtEnd(30)
	// dll.InsertAtEnd(40)
	// dll.InsertAtEnd(50)
	// node := dll.Search(40)
	// fmt.Printf("Searching for 40: Node found with value %d\n", node.Value) // Output: Node found with value 40

	// // Test Contains()
	// fmt.Println("Contains 30:", dll.Contains(30)) // Output: true
	// fmt.Println("Contains 60:", dll.Contains(60)) // Output: false

	// // Test ToArray()
	// arr := dll.ToArray()
	// fmt.Println("Array representation of the list:", arr) // Output: [30 40]
}

// Testing case scenario

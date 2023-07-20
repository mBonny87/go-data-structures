package queue

import (
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	llq := LinkedListQueue{}

	llq.Enqueue(5)
	llq.Enqueue(10)
	llq.Enqueue(15)
	llq.Enqueue(20)
	llq.Dequeue()
	llq.Dequeue()
	if llq.Peek() != 15 {
		t.Errorf("Expected 15 but found %v", llq.Peek())
	}
	if llq.IsEmpty() != false {
		t.Errorf("Expected false but found %v", llq.IsEmpty())
	}
	if llq.Size() != 2 {
		t.Errorf("Expected 2 but found %v", llq.Size())
	}

	llq.Print()
}

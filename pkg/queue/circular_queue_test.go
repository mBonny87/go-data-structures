package queue

import (
	"testing"
)

func TestCircularQueue(t *testing.T) {
	cq := CircularQueue{}

	cq.Enqueue(5)
	cq.Enqueue(10)
	cq.Enqueue(15)
	cq.Enqueue(20)
	cq.Dequeue()
	cq.Dequeue()
	cq.Dequeue()
	if cq.Peek() != 20 {
		t.Errorf("Expected 20 but found %v", cq.Peek())
	}
	if cq.IsEmpty() {
		t.Errorf("Expected false but found %v", cq.IsEmpty())
	}
	if cq.Size() != 1 {
		t.Errorf("Expected 1 but found %v", cq.Size())
	}

	cq.Print() // 5 15
}

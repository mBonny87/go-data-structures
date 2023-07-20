package queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue{}

	pq.Enqueue(5, 35)
	pq.Enqueue(5, 10)
	pq.Enqueue(10, 25)
	pq.Enqueue(15, 10)
	pq.Enqueue(20, 35)
	pq.Dequeue() //dequeue front
	pq.Dequeue() //dequeue rear
	pq.Dequeue() // dequeue middle
	if pq.Peek() != 5 {
		t.Errorf("Expected 5 but found %v", pq.Peek())
	}
	if pq.IsEmpty() != false {
		t.Errorf("Expected false but found %v", pq.IsEmpty())
	}
	if pq.Size() != 2 {
		t.Errorf("Expected 2 but found %v", pq.Size())
	}

	pq.Print() // 5 15
}

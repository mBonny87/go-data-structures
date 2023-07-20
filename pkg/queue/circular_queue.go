package queue

import "fmt"

// That's dynamic, not fixed size
type CircularQueue struct {
	Front *Node
	Rear  *Node
}

func (cq *CircularQueue) Enqueue(v int) {
	if cq.Front == nil {
		cq.Front = &Node{Value: v, Next: nil}
		cq.Rear = cq.Front
		cq.Front.Next = cq.Rear
		cq.Rear.Next = cq.Front
		return
	}
	curr := cq.Front
	for curr.Next != cq.Front {
		curr = curr.Next
	}
	curr.Next = &Node{Value: v, Next: cq.Front}
	cq.Rear = curr.Next
}

func (cq *CircularQueue) Dequeue() {
	if cq.Front == nil {
		//no nodes
		return
	}
	if cq.Front.Next == nil {
		// one node, return empty
		return
	}
	cq.Front = cq.Front.Next
	cq.Rear.Next = cq.Front
}

func (cq *CircularQueue) Peek() int {
	if cq.Front == nil {
		return 0
	}
	return cq.Front.Value
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.Front == nil
}

func (cq *CircularQueue) Size() int {
	if cq.Front == nil {
		return 0
	}
	i := 1
	curr := cq.Front
	for curr != cq.Rear {
		i++
		curr = curr.Next
	}
	return i
}

func (cq *CircularQueue) Clear() {
	cq.Front = nil
	cq.Rear = nil
}

func (cq *CircularQueue) Print() {
	if cq.Front == nil {
		return
	}
	curr := cq.Front
	for curr.Next != cq.Front {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	fmt.Println(curr.Value)
}

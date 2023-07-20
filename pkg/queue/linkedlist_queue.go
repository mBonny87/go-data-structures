package queue

import "fmt"

type LinkedListQueue struct {
	Front *Node
	Rear  *Node
}

func (llq *LinkedListQueue) Enqueue(v int) {
	if llq.Front == nil {
		llq.Front = &Node{Value: v, Next: nil}
		llq.Rear = llq.Front
		return
	}
	curr := llq.Front
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &Node{Value: v, Next: nil}
	llq.Rear = curr.Next
}

func (llq *LinkedListQueue) Dequeue() {
	if llq.Front == nil {
		return
	}
	llq.Front = llq.Front.Next
	if llq.Front == nil {
		llq.Rear = nil
	}
}

func (llq *LinkedListQueue) Peek() int {
	if llq.Front == nil {
		return 0 //means nil
	}
	return llq.Front.Value
}

func (llq *LinkedListQueue) IsEmpty() bool {
	return llq.Front == nil
}

func (llq *LinkedListQueue) Size() int {
	if llq.Front == nil {
		return 0
	}
	i := 1
	curr := llq.Front
	for curr.Next != nil {
		i++
		curr = curr.Next
	}
	return i
}

func (llq *LinkedListQueue) Clear() {
	llq.Front = nil
	llq.Rear = nil
}

func (llq *LinkedListQueue) Print() {
	if llq.Front == nil {
		return
	}
	curr := llq.Front
	for curr.Next != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	// print last value
	fmt.Println(curr.Value)
}

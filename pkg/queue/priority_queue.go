package queue

import "fmt"

type PriorityQueue struct {
	Front *Node
	Rear  *Node
}

//this is a case with an array queue of length 5

func (cq *PriorityQueue) Enqueue(v int, p int) {
	if cq.Front == nil {
		cq.Front = &Node{Value: v, Next: nil, Priority: p}
		cq.Rear = cq.Front
		return
	}
	curr := cq.Front
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &Node{Value: v, Next: nil, Priority: p}
	cq.Rear = curr.Next
}

func (cq *PriorityQueue) Dequeue() {
	if cq.Front == nil {
		return
	}
	curr := cq.Front
	toBeRemoved := curr
	// get the priority node
	for curr.Next != nil {
		if toBeRemoved.Priority < curr.Next.Priority {
			toBeRemoved = curr.Next
		}
		curr = curr.Next
	}
	//if the node to be removed is in the front
	if toBeRemoved == cq.Front {
		cq.Front = cq.Front.Next
		return
	}
	// if the node to be removed is in the rear
	if toBeRemoved == cq.Rear {
		rearCurr := cq.Front
		for rearCurr.Next != nil {
			if rearCurr.Next.Next == nil {
				rearCurr.Next = nil
				cq.Rear = rearCurr
				return
			}
			rearCurr = rearCurr.Next
		}
	}
	// if the node to be removed is in the middle
	middle := cq.Front
	for middle.Next != nil {
		if middle.Next == toBeRemoved {
			middle.Next = middle.Next.Next
			return
		}
		middle = middle.Next
	}

}

func (cq *PriorityQueue) Peek() int {
	if cq.Front == nil {
		return 0
	}
	return cq.Front.Value
}

func (cq *PriorityQueue) IsEmpty() bool {
	return cq.Front == nil
}

func (cq *PriorityQueue) Size() int {
	if cq.Front == nil {
		return 0
	}
	i := 1
	curr := cq.Front
	for curr.Next != nil {
		i++
		curr = curr.Next
	}
	return i
}

func (cq *PriorityQueue) Clear() {
	cq.Front = nil
	cq.Rear = nil
}

func (cq *PriorityQueue) Print() {
	if cq.Front == nil {
		return
	}
	curr := cq.Front
	for curr.Next != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	fmt.Println(curr.Value)
}

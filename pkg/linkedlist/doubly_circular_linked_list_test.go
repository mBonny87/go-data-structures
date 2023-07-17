package linkedlist

import "testing"

func TestDoublyCircularLinkedList(t *testing.T) {
	dcll := DoublyCircularLinkedList{}

	dcll.Push(5)
	dcll.Push(10)
	dcll.Push(15)
	dcll.Push(20)

	dcll.Pop()

	dcll.Unshift(4)
	dcll.Unshift(3)
	dcll.Unshift(2)

	dcll.Shift()
	dcll.Shift()
	dcll.Shift()

	dcll.Print()
}

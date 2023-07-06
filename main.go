package main

import (
	"bonnydev87/datastructures/pkg/linkedlist"
)

func main() {
	dll := linkedlist.CircularLinkedList{}

	dll.Push(10)
	dll.Push(10)
	dll.Push(10)
	dll.Push(10)

	dll.Print()
}

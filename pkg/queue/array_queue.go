package queue

import "fmt"

//this is a case with an array queue of length 5

func Enqueue(v int, queue *[5]int) [5]int {
	// full
	this := (*queue)
	if this[len(this)-1] != 0 {
		return this
	}
	for i, value := range this {
		if value != 0 {
			continue
		} else {
			this[i] = v
			break
		}
	}
	(*queue) = this
	return (*queue)
}

func Dequeue(queue *[5]int) [5]int {
	// check if there's some data
	this := (*queue)
	if this[0] == 0 {
		return this
	}

	//move right indexes to left
	for i := range this {
		if i+1 == len(this) {
			this[i] = 0
			break
		}
		this[i] = this[i+1]
	}
	(*queue) = this
	return (*queue)
}

func Peek(queue *[5]int) int {
	return (*queue)[0]
}

func IsEmpty(queue *[5]int) bool {
	return (*queue)[0] == 0
}

func Size(queue *[5]int) int {
	for i, value := range *queue {
		if value == 0 {
			return i
		}
	}
	return len(queue)
}

func Clear(queue *[5]int) [5]int {
	for i := range *queue {
		(*queue)[i] = 0
	}
	return (*queue)
}

func Print(queue *[5]int) {
	for _, value := range *queue {
		fmt.Println(value)
	}
}

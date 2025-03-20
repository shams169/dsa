package main

import "fmt"

type Queue[T any] struct {
	items []T
	size  int
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, capacity),
		size:  capacity,
	}
}

func (q *Queue[T]) IsFull() bool {
	return len(q.items) == q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Enqueue(item T) {
	if q.IsFull() {
		fmt.Println("Queue is full, cannot enqueue any more items")
		return
	}
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		fmt.Println("Queue is empty, cannot dequeue any more items")
		var zeroValue T
		return zeroValue
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return q.items[0]
}

func (q *Queue[T]) Rear() T {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return q.items[len(q.items)-1]
}

func (q *Queue[T]) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	for i := range len(q.items) {
		fmt.Println(q.items[i])
	}
	fmt.Println("-------------------------")
}

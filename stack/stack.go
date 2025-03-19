package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)

}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		var zeroValue T
		return zeroValue, false
	}
	lastIndex := len(s.items) - 1
	return s.items[lastIndex], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
	}

	lastIndex := len(s.items) - 1
	for i := lastIndex; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
	fmt.Println("------------------")
}

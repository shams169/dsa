package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *SinglyLinkedList) Append(data int) {
	// Prepare a node with the data
	temp := &Node{
		Data: data,
		Next: nil,
	}

	// Check if the list is empty
	if ll.Head == nil {
		ll.Head = temp
		ll.Tail = temp
		return
	}

	ll.Tail.Next = temp
	ll.Tail = temp
}

func (ll *SinglyLinkedList) InsertAt(data, index int) error {

	temp := &Node{
		Data: data,
		Next: nil,
	}

	// Check if the index is greater than or equal to 0
	if index < 0 {
		return fmt.Errorf("index %d is invalid", index)
	}

	//Check if the list is empty and index value is greater than 0
	if index != 0 && ll.Head == nil {
		return fmt.Errorf("list is empty cannot do insert at %d ", index)
	}

	if index == 0 && ll.Head == nil {
		ll.Head = temp
		ll.Tail = temp
		return nil
	}

	if index == 0 && ll.Head != nil {
		temp.Next = ll.Head
		ll.Head = temp
		return nil
	}

	curr := ll.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
		if curr == nil {
			return fmt.Errorf("insufficient number of elements to do insert at index %d", index)
		}
	}

	temp.Next = curr.Next
	curr.Next = temp

	if temp.Next == nil {
		ll.Tail = temp
	}

	return nil
}

func (ll *SinglyLinkedList) Size() int {
	size := 0
	if ll.Head == nil {
		return size
	}
	temp := ll.Head
	for temp != nil {
		temp = temp.Next
		size++
	}
	return size
}

func (ll *SinglyLinkedList) ValueAt(index int) int {
	//Check if the list is empty
	if ll.Head == nil {
		fmt.Printf("list is empty")
		return 0
	}

	// Check if the index is greater than or equal to 0
	if index < 0 {
		fmt.Printf("index %d is invalid", index)
		return 0
	}

	temp := ll.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
		if temp == nil {
			fmt.Printf("index %d is out of range", index)
			return 0
		}
	}
	return temp.Data

}

func (ll *SinglyLinkedList) PushFront(data int) {
	temp := &Node{
		Data: data,
		Next: nil,
	}

	// Check if the list is empty:
	if ll.Head == nil {
		ll.Head = temp
		return
	}
	temp.Next = ll.Head
	ll.Head = temp
}

func (ll *SinglyLinkedList) PushBack(data int) {
	temp := &Node{
		Data: data,
		Next: nil,
	}

	// Check if the list is empty:
	if ll.Head == nil {
		ll.Head = temp
		ll.Tail = temp
		return
	}

	ll.Tail.Next = temp
	ll.Tail = temp
}

func (ll *SinglyLinkedList) Reverse() {
	//Check if the list is empty
	if ll.Head == nil {
		fmt.Println("list is empty")
		return
	}

	// Check if the list has only one element
	if ll.Size() == 1 {
		fmt.Println("List has only one element")
		return
	}

	var prev *Node
	curr := ll.Head
	ll.Tail = ll.Head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	ll.Head = prev

}

func (ll *SinglyLinkedList) PrintSinglyLL() {
	temp := ll.Head
	if temp == nil {
		fmt.Println("The list is empty")
	}
	index := 0
	for temp != nil {
		fmt.Printf("Data at %d: %d\n", index, temp.Data)
		temp = temp.Next
		index++
	}
}

func (ll *SinglyLinkedList) PrintHeadTail() {
	if ll.Head != nil {
		fmt.Printf("Head: %d\n", ll.Head.Data)
	} else {
		fmt.Println("Head: nil")
	}
	if ll.Tail != nil {
		fmt.Printf("Tail: %d\n", ll.Tail.Data)

	} else {
		fmt.Println("Tail: nil")
	}
}

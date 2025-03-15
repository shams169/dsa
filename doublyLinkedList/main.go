package main

func main() {
	// InsertAtHead
	// dll := DoublyLinkedList{}
	// dll.PrintDLL()
	// dll.InsertAtHead(1)
	// dll.PrintDLL()
	// dll.InsertAtHead(2)
	// dll.PrintDLL()
	// dll.InsertAtHead(3)
	// dll.PrintDLL()

	// InsertAtTail
	// dll := DoublyLinkedList{}
	// dll.PrintDLL()
	// dll.InsertAtTail(1)
	// dll.PrintDLL()
	// dll.InsertAtTail(2)
	// dll.PrintDLL()
	// dll.InsertAtTail(3)
	// dll.InsertAtTail(4)
	// dll.InsertAtTail(5)
	// dll.PrintDLL()

	// InsertAtIndex
	dll := DoublyLinkedList{}
	dll.InsertAtIndex(1, 1)
	dll.InsertAtIndex(1, 0)
	dll.PrintDLL()
	dll.InsertAtIndex(2, 0)
	dll.InsertAtIndex(3, 1)

	dll.InsertAtIndex(4, 2)
	dll.InsertAtIndex(5, 3)
	dll.PrintDLL()

	dll.DeleteAtIndex(0)
	dll.PrintDLL()
	dll.DeleteAtIndex(2)
	dll.PrintDLL()

	dll.DeleteAtIndex(1)
	dll.PrintDLL()

	dll.DeleteAtIndex(1)
	dll.PrintDLL()

	dll.DeleteAtIndex(0)
	dll.PrintDLL()

	dll.InsertAtIndex(1, 0)
	dll.PrintDLL()

}

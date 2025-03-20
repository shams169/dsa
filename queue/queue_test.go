package main

import "testing"

func TestEnqueue(t *testing.T) {
	queue := Queue[string]{size: 3}
	queue.Enqueue("Hello")
	queue.Enqueue("World")
	queue.Enqueue("!")

	if len(queue.items) != 3 {
		t.Errorf("Expected queue length to be 3, got %d", len(queue.items))
	}
	if queue.items[0] != "Hello" {
		t.Errorf("Expected first item to be 'Hello', got '%s'", queue.items[0])
	}
	if queue.items[1] != "World" {
		t.Errorf("Expected second item to be 'World', got '%s'", queue.items[1])
	}
	if queue.IsFull() != true {
		t.Errorf("Expected queue to be full, got false")
	}
	queue.Enqueue("One")
	if queue.IsFull() != true && len(queue.items) != 3 {
		t.Errorf("Expected queue to be full, and new item should not get enqueued")
	}
}

func TestDequeue(t *testing.T) {
	queue := Queue[string]{size: 3}
	if queue.IsEmpty() != true {
		t.Errorf("Expected queue to be empty, got false")
	}
	queue.Enqueue("Hello")
	queue.Enqueue("World")
	queue.Enqueue("!")

	// Check the initial size matches the number of items enqueued
	if len(queue.items) != 3 {
		t.Errorf("Expected queue length to be 3, got %d", len(queue.items))
	}

	item := queue.Dequeue()
	if item != "Hello" && len(queue.items) != 2 {
		t.Errorf("Expected first item to be 'Hello' and no of items in the queue to be 2, got '%s' and %d", item, len(queue.items))
	}
	queue.Enqueue("One")
	if queue.IsFull() != true {
		t.Errorf("Expected queue to be full, got false")
	}

}

func TestPeek(t *testing.T) {
	t.Run("Peek on empty queue", func(t *testing.T) {
		queue := NewQueue[string](3)
		item := queue.Peek()
		if item != "" {
			t.Errorf("Expected peek on empty queue to return zero value, got %s", item)
		}
	})

	t.Run("Peek on non-empty queue", func(t *testing.T) {
		queue := NewQueue[string](3)
		queue.Enqueue("Hello")
		queue.Enqueue("World")
		item := queue.Peek()
		if item != "Hello" {
			t.Errorf("Expected peek to return 'Hello', got '%s'", item)
		}
		if len(queue.items) != 2 {
			t.Errorf("Expected queue length to be 2, got %d", len(queue.items))
		}
	})
	t.Run("Peek on non-empty queue with different type", func(t *testing.T) {
		queue := NewQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		item := queue.Peek()
		if item != 1 {
			t.Errorf("Expected peek to return 1, got %d", item)
		}
		if len(queue.items) != 2 {
			t.Errorf("Expected queue length to be 2, got %d", len(queue.items))
		}
	})
}

func TestRear(t *testing.T) {
	t.Run("Rear on empty queue", func(t *testing.T) {
		queue := NewQueue[int](3)
		item := queue.Rear()
		if item != 0 {
			t.Errorf("Expected Rear on empty queue to return zero value, got %d", item)
		}
	})

	t.Run("Rear on non-empty queue", func(t *testing.T) {
		queue := NewQueue[string](3)
		queue.Enqueue("Hello")
		queue.Enqueue("World")
		item := queue.Rear()
		if item != "World" {
			t.Errorf("Expected Rear to return zero value, got '%s'", item)
		}

	})
	t.Run("Rear on non-empty queue with different type", func(t *testing.T) {
		queue := NewQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		item := queue.Rear()
		if item != 3 {
			t.Errorf("Expected Rear to return 0, got %d", item)
		}
	})
	t.Run("Rear on non-empty queue with different type and full queue", func(t *testing.T) {
		queue := NewQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		item := queue.Rear()
		if item != 3 {
			t.Errorf("Expected Rear to return 3, got %d", item)
		}
		if queue.IsFull() != true {
			t.Errorf("Expected queue to be full, got false")
		}
	})
}

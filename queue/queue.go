package queue

import (
	"errors"

	"github.com/grandquista/data-structures-and-algorithms/node"
)

// Queue queue.
type Queue struct {
	Head *node.Node
	size int
}

// MakeQueue Initialize new list with optional iterable.
func MakeQueue(it []interface{}) *Queue {
	queue := Queue{}
	for _, elem := range it {
		queue.Enqueue(elem)
	}
	return &queue
}

// Len Return the number of values currently in the queue.
func (queue Queue) Len() int {
	return queue.size
}

// Dequeue Retrieve and remove the earliest item from the queue.
func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.Head == nil {
		return nil, errors.New("empty queue")
	}
	node := queue.Head
	if node.Next == nil {
		queue.Head = nil
		queue.size -= 1
		return node.Value, nil
	}
	for node.Next.Next != nil {
		node = node.Next
	}
	final := node.Next
	node.Next = nil
	queue.size -= 1
	return final.Value, nil
}

// Enqueue Insert a value into the queue.
func (queue *Queue) Enqueue(value interface{}) {
	queue.Head = node.MakeNode(value, queue.Head)
	queue.size += 1
}

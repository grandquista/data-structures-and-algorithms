package queue

import (
	"errors"

	"github.com/grandquista/data-structures-and-algorithms/node"
)

// Queue queue.
type Queue interface {
	Dequeue() (interface{}, error)
	Enqueue(interface{})
	Len() int
}

type queue struct {
	head *node.Node
	size int
}

// MakeQueue initialize new queue with optional array.
func MakeQueue(it []interface{}) Queue {
	q := queue{}
	for _, elem := range it {
		q.Enqueue(elem)
	}
	return &q
}

// Len Return the number of values currently in the queue.
func (queue queue) Len() int {
	return queue.size
}

// Dequeue Retrieve and remove the earliest item from the queue.
func (queue *queue) Dequeue() (interface{}, error) {
	if queue.head == nil {
		return nil, errors.New("empty queue")
	}
	node := queue.head
	if node.Next == nil {
		queue.head = nil
		queue.size--
		return node.Value, nil
	}
	for node.Next.Next != nil {
		node = node.Next
	}
	final := node.Next
	node.Next = nil
	queue.size--
	return final.Value, nil
}

// Enqueue Insert a value into the queue.
func (queue *queue) Enqueue(value interface{}) {
	queue.head = node.MakeNode(value, queue.head)
	queue.size++
}

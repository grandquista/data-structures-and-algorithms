package queue_test

import (
	"github.com/grandquista/data-structures-and-algorithms/queue"
	"testing"
)

func newQueue() queue.Queue {
	return queue.MakeQueue(nil)
}

func orderedQueue() queue.Queue {
	input := make([]interface{}, 0, 0xFFFFFF)
	for i := 3; i < 40; i += 3 {
		input = append(input, i)
	}
	return queue.MakeQueue(input)
}

func unorderedQueue() queue.Queue {
	input := make([]interface{}, 0, 0xFFFFFF)
	for i := 73; i > 40; i -= 2 {
		input = append(input, i%7)
	}
	return queue.MakeQueue(input)
}

func largeQueue() queue.Queue {
	input := make([]interface{}, 0, 0xFFFFFF)
	for i := 0; i < 0xFFFFFF; i += 1 {
		input = append(input, "task")
	}
	return queue.MakeQueue(input)
}

func TestEmptyQueueDequeue(t *testing.T) {
	queue := newQueue()
	elem, err := queue.Dequeue()
	if err == nil {
		t.Fail()
	}
	if elem != nil {
		t.Fail()
	}
}

func TestEmptyQueueHasSize(t *testing.T) {
	queue := newQueue()
	if queue.Len() != 0 {
		t.Fail()
	}
}

func TestDataQueueDequeueChangesSize(t *testing.T) {
	queue := orderedQueue()
	if queue.Len() != 13 {
		t.Fail()
	}
	switch val, err := queue.Dequeue(); val {
	case 3:
		break
	default:
		t.Error(err)
	}
	if queue.Len() != 12 {
		t.Fail()
	}
}

func TestDataQueueDequeue(t *testing.T) {
	queue := orderedQueue()
	switch val, err := queue.Dequeue(); val {
	case 3:
		break
	default:
		t.Error(err)
	}
	switch val, err := queue.Dequeue(); val {
	case 6:
		break
	default:
		t.Error(err)
	}
}

func TestDataQueueDequeueExaust(t *testing.T) {
	queue := orderedQueue()
	for queue.Len() > 0 {
		queue.Dequeue()
	}
	if queue.Len() != 0 {
		t.Fail()
	}
	elem, err := queue.Dequeue()
	if err == nil {
		t.Fail()
	}
	if elem != nil {
		t.Fail()
	}
}

func TestUnorderedDequeue(t *testing.T) {
	queue := unorderedQueue()
	switch val, err := queue.Dequeue(); val {
	case 3:
		break
	default:
		t.Error(err)
	}
	switch val, err := queue.Dequeue(); val {
	case 1:
		break
	default:
		t.Error(err)
	}
	switch val, err := queue.Dequeue(); val {
	case 6:
		break
	default:
		t.Error(err)
	}
	switch val, err := queue.Dequeue(); val {
	case 4:
		break
	default:
		t.Error(err)
	}
}

func TestEmptyQueueEnqueue(t *testing.T) {
	queue := newQueue()
	queue.Enqueue(0)
	switch val, err := queue.Dequeue(); val {
	case 0:
		break
	default:
		t.Error(err)
	}
}

func TestEmptyQueueEnqueueMultiple(t *testing.T) {
	queue := newQueue()
	for i := 0; i < 30; i++ {
		queue.Enqueue(0)
	}
	queue.Enqueue(1)
	if queue.Len() != 31 {
		t.Fail()
	}
	switch val, err := queue.Dequeue(); val {
	case 0:
		break
	default:
		t.Error(err)
	}
}

func TestEmptyQueueEnqueueChangesSize(t *testing.T) {
	queue := newQueue()
	if queue.Len() != 0 {
		t.Fail()
	}
	queue.Enqueue("test")
	if queue.Len() != 1 {
		t.Fail()
	}
}

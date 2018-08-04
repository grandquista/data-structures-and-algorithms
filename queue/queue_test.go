package queue_test

func newQueue() queue.Queue {
    return queue.Queue{}
}

func orderedQueue() queue.Queue {
  queue := queue.Queue{}
  for i:=3; i<40; i+=3 {
    queue.Enqueue(i)
  }
  return queue
}

func unorderedQueue() queue.Queue {
  queue := queue.Queue{}
  for i:=73; i>40; i-=2 {
    queue.Enqueue(i % 7)
  }
  return queue
}

func largeQueue() {
  input := make([]string, 0, 0xFFFFFF)
  for i:=0; i<0xFFFFFF; i+=1 {
    input = append(input, "task")
  }
  return queue.MakeQueue(input)
}

func TestEmptyQueueDefault(t *testing.T) {
  queue := newQueue()
  if queue != nil {
    t.Fail()
  }
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
  if len(queue) != 0 {
    t.Fail()
  }
}

func TestDataQueueDequeueChangesSize(t *testing.T) {
  queue := orderedQueue()
  if len(orderedQueue) != 13 {
    t.Fail()
  }
  switch queue.Dequeue() {
  case 3:
    break
  default:
    t.Fail()
  }
  if len(orderedQueue) != 12 {
    t.Fail()
  }
}

func TestDataQueueDequeue(t *testing.T) {
  queue := orderedQueue()
  switch queue.Dequeue() {
  case 3:
    break
  default:
    t.Fail()
  }
  switch queue.Dequeue() {
  case 6:
    break
  default:
    t.Fail()
  }
}

func TestDataQueueDequeueExaust(orderedQueue) {
    while orderedQueue {
        orderedQueue.dequeue()
    assert len(orderedQueue) == 0
    with raises(IndexError) {
        orderedQueue.dequeue()


func TestUnorderedDequeue(unorderedQueue) {
    assert unorderedQueue.dequeue() == 3
    assert unorderedQueue.dequeue() == 1
    assert unorderedQueue.dequeue() == 6
    assert unorderedQueue.dequeue() == 4


func TestEmptyQueueEnqueue(newQueue) {
    newQueue.enqueue(0)
    assert newQueue.head.value == 0


func TestEmptyQueueEnqueueMultiple(newQueue) {
    for _ in range(30) {
        newQueue.enqueue(0)
    newQueue.enqueue(1)
    assert len(newQueue) == 31
    assert newQueue.dequeue() == 0


func TestEmptyQueueEnqueueChangesSize(newQueue) {
    assert len(newQueue) == 0
    newQueue.enqueue("test")
    assert len(newQueue) == 1

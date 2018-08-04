from pytest import raises

from .queue_with_stacks import Queue
from pytest import fixture


@fixture
def new_queue():
    return Queue()


@fixture
def ordered_queue():
    queue = Queue()
    set(map(queue.enqueue, range(3, 40, 3)))
    return queue


@fixture
def unordered_queue():
    queue = Queue()
    set(map(queue.enqueue, map(lambda i: i % 7, range(73, 40, -2))))
    return queue


@fixture
def large_queue():
    queue = Queue()
    set(map(queue.enqueue, range(0xFFFFFF)))
    return queue

def test_empty_queue_dequeue(new_queue):
    with raises(IndexError):
        new_queue.dequeue()


def test_empty_queue_has_size(new_queue):
    assert len(new_queue) == 0


def test_data_queue_dequeue_changes_size(ordered_queue):
    assert len(ordered_queue) == 13
    assert ordered_queue.dequeue() == 3
    assert len(ordered_queue) == 12


def test_data_queue_dequeue(ordered_queue):
    assert ordered_queue.dequeue() == 3
    assert ordered_queue.dequeue() == 6


def test_data_queue_dequeue_exaust(ordered_queue):
    while ordered_queue:
        ordered_queue.dequeue()
    assert len(ordered_queue) == 0
    with raises(IndexError):
        ordered_queue.dequeue()


def test_unordered_dequeue(unordered_queue):
    assert unordered_queue.dequeue() == 3
    assert unordered_queue.dequeue() == 1
    assert unordered_queue.dequeue() == 6
    assert unordered_queue.dequeue() == 4


def test_empty_queue_enqueue(new_queue):
    new_queue.enqueue(0)
    assert new_queue.dequeue() == 0


def test_empty_queue_enqueue_multiple(new_queue):
    for _ in range(30):
        new_queue.enqueue(0)
    new_queue.enqueue(1)
    assert len(new_queue) == 31
    assert new_queue.dequeue() == 0


def test_empty_queue_enqueue_changes_size(new_queue):
    assert len(new_queue) == 0
    new_queue.enqueue("test")
    assert len(new_queue) == 1

from pytest import raises

from .stack import Stack
from pytest import fixture


@fixture
def new_stack():
    return Stack()


@fixture
def ordered_stack():
    return Stack(range(3, 40, 3))


@fixture
def unordered_stack():
    return Stack(map(lambda i: i % 7, range(73, 40, -2)))


@fixture
def large_stack():
    return Stack("task" for _ in range(0xFFFFFF))

def test_empty_stack_default(new_stack):
    assert new_stack.head is None


def test_empty_stack_pop(new_stack):
    with raises(IndexError):
        new_stack.pop()


def test_empty_stack_peek(new_stack):
    with raises(IndexError):
        new_stack.peek()


def test_empty_stack_has_size(new_stack):
    assert len(new_stack) == 0


def test_data_stack_pop_changes_size(ordered_stack):
    assert len(ordered_stack) == 13
    assert ordered_stack.pop() == 39
    assert len(ordered_stack) == 12


def test_data_stack_peek_no_mutate(ordered_stack):
    assert len(ordered_stack) == 13
    assert ordered_stack.peek() == 39
    assert len(ordered_stack) == 13
    assert ordered_stack.peek() == 39


def test_data_stack_pop(ordered_stack):
    assert ordered_stack.pop() == 39
    assert ordered_stack.pop() == 36


def test_data_stack_pop_exaust(ordered_stack):
    while ordered_stack:
        ordered_stack.pop()
    assert len(ordered_stack) == 0
    with raises(IndexError):
        ordered_stack.pop()


def test_unordered_pop(unordered_stack):
    assert unordered_stack.pop() == 6
    assert unordered_stack.pop() == 1
    assert unordered_stack.pop() == 3
    assert unordered_stack.pop() == 5


def test_empty_stack_push(new_stack):
    new_stack.push(0)
    assert new_stack.head.value == 0


def test_empty_stack_push_multiple(new_stack):
    for _ in range(30):
        new_stack.push(0)
    new_stack.push(1)
    assert len(new_stack) == 31
    assert new_stack.pop() == 1


def test_empty_stack_push_changes_size(new_stack):
    assert len(new_stack) == 0
    new_stack.push("test")
    assert len(new_stack) == 1

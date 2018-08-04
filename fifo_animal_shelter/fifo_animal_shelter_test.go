from .fifo_animal_shelter import Cat, Dog
from pytest import raises

from .fifo_animal_shelter import AnimalShelter, Cat, Dog
from pytest import fixture


@fixture
def new_queue():
    return AnimalShelter()


@fixture
def ordered_queue():
    return AnimalShelter(Dog() if i < 10 else Cat() for i in range(3, 40, 3))


@fixture
def unordered_queue():
    return AnimalShelter(
        Dog() if i % 7 < 3 else Cat() for i in range(73, 40, -2)
    )


@fixture
def large_queue():
    return AnimalShelter(Cat() for _ in range(0xFFFFFF))

def test_empty_queue_dequeue(new_queue):
    with raises(IndexError):
        new_queue.dequeue()


def test_content_in_empty_queue_str(new_queue):
    assert len(str(new_queue)) > 0


def test_content_in_queue_str(ordered_queue):
    assert len(str(ordered_queue)) > 0


def test_content_in_empty_queue_repr(new_queue):
    assert len(repr(new_queue)) > 0


def test_content_in_queue_repr(unordered_queue):
    assert len(repr(unordered_queue)) > 0


def test_empty_queue_has_size(new_queue):
    assert len(new_queue) == 0


def test_data_queue_dequeue_changes_size(ordered_queue):
    assert len(ordered_queue) == 13
    assert isinstance(ordered_queue.dequeue(), Dog)
    assert len(ordered_queue) == 12


def test_data_queue_dequeue(ordered_queue):
    assert isinstance(ordered_queue.dequeue(), Dog)
    assert isinstance(ordered_queue.dequeue(), Dog)


def test_data_queue_dequeue_exaust(ordered_queue):
    while ordered_queue:
        ordered_queue.dequeue()
    assert len(ordered_queue) == 0
    with raises(IndexError):
        ordered_queue.dequeue()


def test_unordered_dequeue(unordered_queue):
    assert isinstance(unordered_queue.dequeue(), Cat)
    assert isinstance(unordered_queue.dequeue(), Dog)
    assert isinstance(unordered_queue.dequeue(), Cat)
    assert isinstance(unordered_queue.dequeue(), Cat)


def test_enqueue_not_animal_raises(new_queue):
    with raises(TypeError):
        new_queue.enqueue(None)
    with raises(TypeError):
        new_queue.enqueue("None")
    with raises(TypeError):
        new_queue.enqueue(3.14)


def test_empty_queue_enqueue_multiple(new_queue):
    for _ in range(30):
        new_queue.enqueue(Cat())
    new_queue.enqueue(Dog())
    assert len(new_queue) == 31
    assert isinstance(new_queue.dequeue(), Cat)
    assert isinstance(new_queue.dequeue(Dog), Dog)


def test_empty_queue_dequeue_spare(new_queue):
    for _ in range(30):
        new_queue.enqueue(Cat())
    new_queue.enqueue(Dog())
    assert len(new_queue) == 31
    assert isinstance(new_queue.dequeue(Dog), Dog)
    assert isinstance(new_queue.dequeue(), Cat)


def test_empty_queue_dequeue_spare_late(new_queue):
    for _ in range(30):
        new_queue.enqueue(Cat())
    for _ in range(15):
        new_queue.enqueue(Dog())
    assert isinstance(new_queue.dequeue(Dog), Dog)
    assert isinstance(new_queue.dequeue(Cat), Cat)
    assert isinstance(new_queue.dequeue(), Cat)


def test_empty_queue_enqueue_changes_size(new_queue):
    assert len(new_queue) == 0
    new_queue.enqueue(Cat())
    assert len(new_queue) == 1


def test_dequeue_not_animal_raises(unordered_queue):
    with raises(ValueError):
        unordered_queue.dequeue(str)
    with raises(ValueError):
        unordered_queue.dequeue(int)

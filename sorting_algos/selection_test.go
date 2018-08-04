from random import shuffle

from .selection import selection


def selection_test(input):
    output = selection(input)
    return output, sorted(output)


def test_empty_mergesort():
    left, right = selection_test([])
    assert left == right


def test_small_mergesort():
    left, right = selection_test([5, 3, 1, 4, 2])
    assert left == right


def test_invert_mergesort():
    left, right = selection_test(list(range(10, 0, -1)))
    assert left == right


def test_large_mergesort():
    left, right = selection_test([i for _ in range(1000) for i in range(3)])
    assert left == right


def test_sorted_mergesort():
    left, right = selection_test(
        [i for j in range(100) for i in range(j, j + 2)]
    )
    assert left == right


def test_random_mergesort():
    x = list(range(1000))
    for _ in range(100):
        shuffle(x)
        left, right = selection_test(x)
        assert left == right

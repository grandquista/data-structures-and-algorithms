from random import shuffle

from .quicksort import quicksort


def quicksort_test(input):
    """
    Helper to compare sort methods.
    """
    output = quicksort(input)
    return output, sorted(output)


def test_empty_quicksort():
    """
    Test empty array with quicksort.
    """
    left, right = quicksort_test([])
    assert left == right


def test_small_quicksort():
    """
    Test small array with quicksort.
    """
    left, right = quicksort_test([5, 3, 1, 4, 2])
    assert left == right


def test_invert_quicksort():
    """
    Test inverted array with quicksort.
    """
    left, right = quicksort_test(list(range(10, 0, -1)))
    assert left == right


def test_same_quicksort():
    """
    Test inverted array with quicksort.
    """
    left, right = quicksort_test([2, 1, 2] + [0] * 10 + [2, 1, 2])
    assert left == right


def test_large_quicksort():
    """
    Test large array with quicksort.
    """
    left, right = quicksort_test([i for _ in range(1000) for i in range(3)])
    assert left == right


def test_sorted_quicksort():
    """
    Test sorted array with quicksort.
    """
    left, right = quicksort_test(
        [i for j in range(100) for i in range(j, j + 2)]
    )
    assert left == right


def test_random_quicksort():
    """
    Test random array with quicksort.
    """
    x = list(range(1000))
    for _ in range(100):
        shuffle(x)
        left, right = quicksort_test(x)
        assert left == right

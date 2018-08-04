from random import shuffle

from .radix_sort import radix_sort


def radix_sort_test(input):
    """
    Helper to compare sort methods.
    """
    output = radix_sort(input)
    return output, sorted(output)


def test_empty_radix_sort():
    """
    Test empty array with radix_sort.
    """
    left, right = radix_sort_test([])
    assert left == right


def test_small_radix_sort():
    """
    Test small array with radix_sort.
    """
    left, right = radix_sort_test([5, 3, 1, 4, 2])
    assert left == right


def test_scaling_numbers_radix_sort():
    """
    Test scaling numbers array with radix_sort.
    """
    left, right = radix_sort_test([2222, 222, 22, 2])
    assert left == right


def test_spread_radix_sort():
    """
    Test spread numbers array with radix_sort.
    """
    left, right = radix_sort_test([0xffff, 0xfffe, 0xfffd, 3, 2, 1])
    assert left == right


def test_negative_radix_sort():
    """
    Test small array with radix_sort.
    """
    left, right = radix_sort_test([-5, 3, 1, -4, 2])
    assert left == right


def test_shifted_digit_radix_sort():
    """
    Test shifted digit array with radix_sort.
    """
    left, right = radix_sort_test([i << 4 for i in range(100)])
    assert left == right


def test_pow_digit_radix_sort():
    """
    Test power of ten digits array with radix_sort.
    """
    left, right = radix_sort_test([i * 1000 for i in range(100)])
    assert left == right


def test_invert_radix_sort():
    """
    Test inverted array with radix_sort.
    """
    left, right = radix_sort_test(list(range(10, 0, -1)))
    assert left == right


def test_same_radix_sort():
    """
    Test inverted array with radix_sort.
    """
    left, right = radix_sort_test([2, 1, 2] + [0] * 10 + [2, 1, 2])
    assert left == right


def test_large_radix_sort():
    """
    Test large array with radix_sort.
    """
    left, right = radix_sort_test([i for _ in range(1000) for i in range(3)])
    assert left == right


def test_sorted_radix_sort():
    """
    Test sorted array with radix_sort.
    """
    left, right = radix_sort_test(
        [i for j in range(100) for i in range(j, j + 2)]
    )
    assert left == right


def test_random_radix_sort():
    """
    Test random array with radix_sort.
    """
    x = list(range(1000))
    for _ in range(100):
        shuffle(x)
        left, right = radix_sort_test(x)
        assert left == right

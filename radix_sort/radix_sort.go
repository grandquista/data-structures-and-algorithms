from itertools import chain
from math import log


class _IsSorted:
    pass


def _is_sorted(status):
    status.is_sorted = True
    status.max_digit = 0
    last_elem = None

    def _each(elem):
        nonlocal status, last_elem
        if status.is_sorted and last_elem is not None:
            status.is_sorted = last_elem <= elem
        abs_elem = elem
        if status.max_digit < abs_elem:
            status.max_digit = abs_elem
        last_elem = elem
        return elem

    return _each


def _radix_help(a_buckets, b_buckets, digit, radix):
    for bucket in b_buckets:
        bucket.clear()
    for i in chain.from_iterable(a_buckets):
        bucket = (i // (radix ** digit)) % radix
        if i >= 0:
            bucket += radix
        b_buckets[bucket].append(i)
    return b_buckets, a_buckets


def radix_sort(array):
    """
    Sort an array using a radix sort.
    """
    status = _IsSorted()
    array = list(map(_is_sorted(status), array))
    if status.is_sorted:
        return array
    digit_count = status.max_digit.bit_length()
    radix = 1 << (digit_count // 10) + 1
    a_buckets, b_buckets = (
        [[] for _ in range(radix * 2)],
        [[] for _ in range(radix * 2)],
    )
    a_buckets[0] = array
    for digit in range(11):
        a_buckets, b_buckets = _radix_help(a_buckets, b_buckets, digit, radix)
    return list(chain.from_iterable(a_buckets))

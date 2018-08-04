from itertools import chain, islice, repeat, tee


def sort_range(array, start, end):
    """
    Sort an array slice using an iterative mergesort.
    """
    mid = (start + end) // 2
    ks = chain(range(mid, end), repeat(end))
    k = next(ks)
    for j in range(start, mid + 1):
        while array[k] < array[j]:
            array[k], array[j] = array[j], array[k]
            k = next(ks)


def mergesort(array):
    """
    Sort an array using an iterative mergesort.
    """
    array = list(array)
    if not array:
        return array
    len_array = len(array)
    step = 1
    while step < len_array:
        step <<= 1
        steps = tee(chain(range(0, len_array, step), [len_array]))
        odd = islice(steps[0], 0, None)
        even = islice(steps[1], 1, None)
        for start, end in zip(odd, even):
            sort_range(array, start, end)
    sort_range(array, 0, len_array)
    return array

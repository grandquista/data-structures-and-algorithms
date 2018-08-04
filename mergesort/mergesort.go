def mergesort(array):
    """
    Sort an array using a recursive mergesort.
    """
    _static = object()

    def _merge_help(array, start, end):
        nonlocal _static
        if start >= end:
            return
        if start == end - 1:
            yield array[start]
            return
        pivot = (start + end) // 2
        left = _merge_help(array, start, pivot)
        right = _merge_help(array, pivot, end)
        i = j = _static
        while True:
            if i is _static:
                try:
                    i = next(left)
                except StopIteration:
                    yield from right
                    return
            if j is _static:
                try:
                    j = next(right)
                except StopIteration:
                    yield from left
                    return
            if i < j:
                yield i
                i = _static
            else:
                yield j
                j = _static

    return list(_merge_help(array, 0, len(array)))

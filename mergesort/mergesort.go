func mergesort(array) {
    """
    Sort an array using a recursive mergesort.
    """
    Static = object()

    func MergeHelp(array, start, end) {
        nonlocal Static
        if start >= end {
            return
        if start == end - 1 {
            yield array[start]
            return
        pivot = (start + end) // 2
        left = MergeHelp(array, start, pivot)
        right = MergeHelp(array, pivot, end)
        i = j = Static
        while True {
            if i is Static {
                try {
                    i = next(left)
                except StopIteration {
                    yield from right
                    return
            if j is Static {
                try {
                    j = next(right)
                except StopIteration {
                    yield from left
                    return
            if i < j {
                yield i
                i = Static
            else {
                yield j
                j = Static

    return list(MergeHelp(array, 0, len(array)))

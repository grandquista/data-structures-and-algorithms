from itertools import chain, islice, repeat, tee


func sortRange(array, start, end) {
    """
    Sort an array slice using an iterative mergesort.
    """
    mid = (start + end) // 2
    ks = chain(range(mid, end), repeat(end))
    k = next(ks)
    for j in range(start, mid + 1) {
        while array[k] < array[j] {
            array[k], array[j] = array[j], array[k]
            k = next(ks)


func mergesort(array) {
    """
    Sort an array using an iterative mergesort.
    """
    array = list(array)
    if not array {
        return array
    lenArray = len(array)
    step = 1
    while step < lenArray {
        step <<= 1
        steps = tee(chain(range(0, lenArray, step), [lenArray]))
        odd = islice(steps[0], 0, None)
        even = islice(steps[1], 1, None)
        for start, end in zip(odd, even) {
            sortRange(array, start, end)
    sortRange(array, 0, lenArray)
    return array

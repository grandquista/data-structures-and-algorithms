def binary_search(array, key):
    bound_high = len(array) - 1
    bound_low = 0
    index = 0
    while bound_low <= bound_high:
        index = bound_low + (bound_high - bound_low) // 2
        if array[index] == key:
            return index
        elif array[index] < key:
            bound_low = index + 1
        else:
            bound_high = index - 1
    return -1

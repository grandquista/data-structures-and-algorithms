from itertools import islice


def selection(array):
    """
    Sort an array using a selection sort.
    """
    array = list(array)
    for i in range(len(array)):
        swap = min(enumerate(islice(array, i, len(array))), key=lambda t: t[1])
        j = swap[0] + i
        array[i], array[j] = array[j], array[i]
    return array

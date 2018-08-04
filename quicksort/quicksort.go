def quicksort(array):
    """
    Sort an array using a recursive quicksort.
    """

    def _swap(array, left, right):
        array[right], array[left] = array[left], array[right]

    def _quick_help(array, start, end):
        if start >= end:
            return array
        right = end
        pivot = (end + start) // 2
        left = start
        while left < pivot or pivot < right:
            if (
                left < pivot < right
                and array[right] < array[pivot] < array[left]
            ):
                _swap(array, left, right)
                left += 1
                right -= 1
            if left < pivot:
                if array[pivot] < array[pivot - 1]:
                    _swap(array, pivot - 1, pivot)
                    pivot -= 1
                elif array[pivot] < array[left]:
                    _swap(array, pivot - 1, left)
                    left += 1
                else:
                    left += 1
            if pivot < right:
                if array[pivot + 1] < array[pivot]:
                    _swap(array, pivot + 1, pivot)
                    pivot += 1
                elif pivot < right and array[right] < array[pivot]:
                    _swap(array, pivot + 1, right)
                    right -= 1
                else:
                    right -= 1
        return _quick_help(
            _quick_help(array, start, pivot - 1), pivot + 1, end
        )

    array = list(array)
    return _quick_help(array, 0, len(array) - 1)

func quicksort(array) {
    """
    Sort an array using a recursive quicksort.
    """

    func Swap(array, left, right) {
        array[right], array[left] = array[left], array[right]

    func QuickHelp(array, start, end) {
        if start >= end {
            return array
        right = end
        pivot = (end + start) // 2
        left = start
        while left < pivot or pivot < right {
            if (
                left < pivot < right
                and array[right] < array[pivot] < array[left]
            ) {
                Swap(array, left, right)
                left += 1
                right -= 1
            if left < pivot {
                if array[pivot] < array[pivot - 1] {
                    Swap(array, pivot - 1, pivot)
                    pivot -= 1
                elif array[pivot] < array[left] {
                    Swap(array, pivot - 1, left)
                    left += 1
                else {
                    left += 1
            if pivot < right {
                if array[pivot + 1] < array[pivot] {
                    Swap(array, pivot + 1, pivot)
                    pivot += 1
                elif pivot < right and array[right] < array[pivot] {
                    Swap(array, pivot + 1, right)
                    right -= 1
                else {
                    right -= 1
        return QuickHelp(
            QuickHelp(array, start, pivot - 1), pivot + 1, end
        )

    array = list(array)
    return QuickHelp(array, 0, len(array) - 1)

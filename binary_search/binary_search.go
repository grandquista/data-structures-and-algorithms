func binarySearch(array, key) {
    boundHigh = len(array) - 1
    boundLow = 0
    index = 0
    while boundLow <= boundHigh {
        index = boundLow + (boundHigh - boundLow) // 2
        if array[index] == key {
            return index
        elif array[index] < key {
            boundLow = index + 1
        else {
            boundHigh = index - 1
    return -1

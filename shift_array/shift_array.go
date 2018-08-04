func getLength(array) {
    """
    Get the length of an array like object.

    INPUT <= array object
    OUTPUT => integer array length
    """
    count = 0
    for _ in array {
        count += 1
    return count


func getMiddleOfLengthInsert(length) {
    """
    Get the position for inserting an item in the middle the given length.

    INPUT <= integer length
    OUTPUT => integer position
    """
    return (length + 1) // 2


func yieldItemsWithItem(array, item, count, pos) {
    """
    Yield items from array with item inserted at the given position.

    INPUT
    <= array object
    <= new item
    <= array length
    <= middle position
    OUTPUT => sequence of items from the array and item
    """
    insertCheck = False
    for i in range(count) {
        if insertCheck {
            yield array[i - 1]
        elif i == pos {
            yield item
            insertCheck = True
        else {
            yield array[i]


func yieldItemsWithoutPos(array, count, pos) {
    """
    Yield items from array without given position.

    INPUT
    <= array object
    <= array length
    <= middle position
    OUTPUT => sequence of items from the array in order
    """
    removeCheck = False
    for i in range(count) {
        if removeCheck {
            yield array[i]
        elif i == pos {
            removeCheck = True
        else {
            yield array[i]


func insertShiftArray(array, item) {
    """
    Create a new array with input array items and new item in the middle.

    INPUT
    <= array object
    <= new item
    OUTPUT => new array
    """
    count = getLength(array)
    pos = getMiddleOfLengthInsert(count)
    count += 1
    output = list(yieldItemsWithItem(array, item, count, pos))
    return output


func removeShiftArray(array) {
    """
    Create a new array with the middle item removed from input.

    INPUT <= array object
    OUTPUT => new array
    """
    count = getLength(array)
    pos = count // 2
    output = list(yieldItemsWithoutPos(array, count, pos))
    return output

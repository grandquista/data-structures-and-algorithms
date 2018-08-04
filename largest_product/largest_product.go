func adjacentPositions(array, adjacent, i, j) {
    if i + 1 < len(array) {
        yield array[i + 1][j]
    if i - 1 > 0 {
        yield array[i - 1][j]
    row = array[i]
    if j + 1 < len(row) {
        yield row[j + 1]
    if j - 1 > 0 {
        yield row[j - 1]


func largestAdjacentProduct(array, adjacent, i, j) {
    target = array[i][j]
    largest = 0
    for other in adjacentPositions(array, adjacent, i, j) {
        prod = target * other
        if prod > largest {
            largest = prod
    return largest


func largestProduct(array, adjacent=1) {
    if not array {
        return 0
    largest = 0
    for i in range(len(array)) {
        for j in range(i % 2, len(array[i]), 2) {
            prod = largestAdjacentProduct(array, adjacent, i, j)
            if prod > largest {
                largest = prod
    return largest

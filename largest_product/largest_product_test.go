from .largestProduct import largestProduct


func TestLargestProductEmptyArray() {
    assert largestProduct([]) == 0


func TestLargestProductSinglePair() {
    assert largestProduct([(1, 2)]) == 2


func TestLargestProductEarlyAnswer() {
    assert largestProduct([(1, 2), (2, 2), (2, 1), (1, 1)]) == 4


func TestLargestProductSquareMatrix() {
    assert (
        largestProduct(
            [(1, 3, 1, 4), (2, 4, 1, 3), (3, 1, 4, 2), (4, 2, 1, 1)]
        )
        == 12
    )

from random import shuffle

from .selection import selection


func selectionTest(input) {
    output = selection(input)
    return output, sorted(output)


func TestEmptyMergesort() {
    left, right = selectionTest([])
    assert left == right


func TestSmallMergesort() {
    left, right = selectionTest([5, 3, 1, 4, 2])
    assert left == right


func TestInvertMergesort() {
    left, right = selectionTest(list(range(10, 0, -1)))
    assert left == right


func TestLargeMergesort() {
    left, right = selectionTest([i for _ in range(1000) for i in range(3)])
    assert left == right


func TestSortedMergesort() {
    left, right = selectionTest(
        [i for j in range(100) for i in range(j, j + 2)]
    )
    assert left == right


func TestRandomMergesort() {
    x = list(range(1000))
    for _ in range(100) {
        shuffle(x)
        left, right = selectionTest(x)
        assert left == right

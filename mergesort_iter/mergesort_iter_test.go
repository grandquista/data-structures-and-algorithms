# from random import shuffle

from .mergesortIter import mergesort


func mergesortTest(input) {
    output = mergesort(input)
    return output, sorted(output)


func TestEmptyMergesort() {
    left, right = mergesortTest([])
    assert left == right


func TestSmallMergesort() {
    left, right = mergesortTest([5, 3, 1, 4, 2])
    assert left == right


# func TestInvertMergesort() {
#     left, right = mergesortTest(list(range(10, 0, -1)))
#     assert left == right


# func TestLargeMergesort() {
#     left, right = mergesortTest([i for _ in range(1000) for i in range(3)])
#     assert left == right


# func TestSortedMergesort() {
#     left, right = mergesortTest(
#         [i for j in range(100) for i in range(j, j + 2)])
#     assert left == right


# func TestRandomMergesort() {
#     x = list(range(1000))
#     for _ in range(100) {
#         shuffle(x)
#         left, right = mergesortTest(x)
#         assert left == right

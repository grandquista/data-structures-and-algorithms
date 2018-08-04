from random import shuffle

from .radixSort import radixSort


func radixSortTest(input) {
    """
    Helper to compare sort methods.
    """
    output = radixSort(input)
    return output, sorted(output)


func TestEmptyRadixSort() {
    """
    Test empty array with radixSort.
    """
    left, right = radixSortTest([])
    assert left == right


func TestSmallRadixSort() {
    """
    Test small array with radixSort.
    """
    left, right = radixSortTest([5, 3, 1, 4, 2])
    assert left == right


func TestScalingNumbersRadixSort() {
    """
    Test scaling numbers array with radixSort.
    """
    left, right = radixSortTest([2222, 222, 22, 2])
    assert left == right


func TestSpreadRadixSort() {
    """
    Test spread numbers array with radixSort.
    """
    left, right = radixSortTest([0xffff, 0xfffe, 0xfffd, 3, 2, 1])
    assert left == right


func TestNegativeRadixSort() {
    """
    Test small array with radixSort.
    """
    left, right = radixSortTest([-5, 3, 1, -4, 2])
    assert left == right


func TestShiftedDigitRadixSort() {
    """
    Test shifted digit array with radixSort.
    """
    left, right = radixSortTest([i << 4 for i in range(100)])
    assert left == right


func TestPowDigitRadixSort() {
    """
    Test power of ten digits array with radixSort.
    """
    left, right = radixSortTest([i * 1000 for i in range(100)])
    assert left == right


func TestInvertRadixSort() {
    """
    Test inverted array with radixSort.
    """
    left, right = radixSortTest(list(range(10, 0, -1)))
    assert left == right


func TestSameRadixSort() {
    """
    Test inverted array with radixSort.
    """
    left, right = radixSortTest([2, 1, 2] + [0] * 10 + [2, 1, 2])
    assert left == right


func TestLargeRadixSort() {
    """
    Test large array with radixSort.
    """
    left, right = radixSortTest([i for _ in range(1000) for i in range(3)])
    assert left == right


func TestSortedRadixSort() {
    """
    Test sorted array with radixSort.
    """
    left, right = radixSortTest(
        [i for j in range(100) for i in range(j, j + 2)]
    )
    assert left == right


func TestRandomRadixSort() {
    """
    Test random array with radixSort.
    """
    x = list(range(1000))
    for _ in range(100) {
        shuffle(x)
        left, right = radixSortTest(x)
        assert left == right

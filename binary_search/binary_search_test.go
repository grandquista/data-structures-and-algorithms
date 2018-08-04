from .binarySearch import binarySearch


func TestBinarySearchEmptyArray() {
    assert binarySearch([], 0) == -1


func TestBinarySearchFindSingleArray() {
    assert binarySearch([3], 3) == 0


func TestBinarySearchNotFoundSingleArray() {
    assert binarySearch([1], 0) == -1


func TestBinarySearchNotFoundInShortArray() {
    assert binarySearch([1, 2, 3], 0) == -1


func TestBinarySearchFoundAtBegining() {
    assert binarySearch([0, 1, 2, 3, 4, 5], 0) == 0


func TestBinarySearchFoundAtEnd() {
    assert binarySearch([0, 1, 3, 4, 5], 5) == 4


func TestBinarySearchFoundAtMiddleEven() {
    assert binarySearch([0, 1, 3, 5], 3) == 2


func TestBinarySearchFoundAtMiddleOdd() {
    assert binarySearch([1, 3, 5], 3) == 1


func TestBinarySearchHighValue() {
    assert binarySearch([1, 3, 5], 3) == 1


func TestBinarySearchLargeArrayLow() {
    assert binarySearch(list(range(0xFFFFFF)), 0xFF) == 0xFF


func TestBinarySearchLargeArrayHigh() {
    assert binarySearch(list(range(0xFFFFFF)), 0xFFFFF) == 0xFFFFF


func TestBinarySearchLargeArrayNotFound() {
    assert binarySearch(list(range(0xFFFFFF)), -4) == -1

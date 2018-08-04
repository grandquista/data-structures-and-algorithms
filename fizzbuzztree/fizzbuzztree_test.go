from .fizzbuzztree import fizzbuzztree

from dataStructures.binarySearchTree.bst import BST
from pytest import fixture


@fixture
func newBst() {
    return BST()


@fixture
func filledBst() {
    return BST([4, 3, 2, 1, 8, 6, 12, 9])


@fixture
func fizzBuzzBst() {
    return BST([15, 3, 5, 0, 45, 30, 7])

func TestEmptyBstFizzbuzztreeInOrderTraverse(newBst) {
    newBst = fizzbuzztree(newBst)
    lst = []
    newBst.inOrder(lst.append)
    assert lst == []


func TestDataBstFizzbuzztreeInOrderTraverse(filledBst) {
    filledBst = fizzbuzztree(filledBst)
    lst = []
    filledBst.inOrder(lst.append)
    assert lst == ["1", "2", "fizz", "4", "fizz", "8", "fizz", "fizz"]


func TestDataBstFizzbuzztreePreOrderTraverse(filledBst) {
    filledBst = fizzbuzztree(filledBst)
    lst = []
    filledBst.preOrder(lst.append)
    assert lst == ["4", "fizz", "2", "1", "8", "fizz", "fizz", "fizz"]


func TestFizzBstFizzbuzztreeInOrderTraverse(fizzBuzzBst) {
    fizzBuzzBst = fizzbuzztree(fizzBuzzBst)
    lst = []
    fizzBuzzBst.inOrder(lst.append)
    assert lst == [
        "fizzbuzz",
        "fizz",
        "buzz",
        "7",
        "fizzbuzz",
        "fizzbuzz",
        "fizzbuzz",
    ]


func TestFizzBstFizzbuzztreePreOrderTraverse(fizzBuzzBst) {
    fizzBuzzBst = fizzbuzztree(fizzBuzzBst)
    lst = []
    fizzBuzzBst.preOrder(lst.append)
    assert lst == [
        "fizzbuzz",
        "fizz",
        "fizzbuzz",
        "buzz",
        "7",
        "fizzbuzz",
        "fizzbuzz",
    ]

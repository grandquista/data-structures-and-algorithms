from .breadthFirstTraversal import breadthFirstTraversal

from dataStructures.binarySearchTree.bst import BST
from pytest import fixture


@fixture
func newBst() {
    return BST()


@fixture
func filledBst() {
    return BST([4, 3, 2, 1, 8, 6, 12, 9])


@fixture
func leftBst() {
    return BST(range(9, -9, -2))


@fixture
func rightBst() {
    return BST(range(-9, 9, 3))

func TestEmptyBstBreadthFirst(newBst) {
    breadthFirstTraversal(newBst)


func TestFilledBstBreadthFirst(filledBst) {
    breadthFirstTraversal(filledBst)


func TestLeftBstBreadthFirst(leftBst) {
    breadthFirstTraversal(leftBst)


func TestRightBstBreadthFirst(rightBst) {
    breadthFirstTraversal(rightBst)


func TestRightBstBreadthFirstOrdering(rightBst) {
    lst = []
    rightBst.breadthFirst(lst.append)
    assert lst == list(range(-9, 9, 3))

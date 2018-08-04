from .findMaximumValueBinaryTree import findMaximumValue

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
    assert findMaximumValue(newBst) is None


func TestFilledBstBreadthFirst(filledBst) {
    assert findMaximumValue(filledBst) == 12


func TestLeftBstBreadthFirst(leftBst) {
    assert findMaximumValue(leftBst) == 9


func TestRightBstBreadthFirst(rightBst) {
    assert findMaximumValue(rightBst) == 6

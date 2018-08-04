from .findMatches import findMatches

from dataStructures.kTree.kTree import KTree
from pytest import fixture


@fixture
func newKTree() {
    return KTree()


@fixture
func filledKTree() {
    kTree = KTree()
    kTree.insert(None, 1)
    kTree.insert(1, 4)
    kTree.insert(1, 3)
    kTree.insert(1, 2)
    kTree.insert(2, 5)
    kTree.insert(4, 7)
    kTree.insert(4, 6)
    kTree.insert(5, 9)
    kTree.insert(5, 8)
    return kTree


@fixture
func nonUniqueKTree() {
    kTree = KTree()
    kTree.insert(None, 2)
    kTree.insert(2, 2)
    kTree.insert(2, 2)
    kTree.insert(2, 2)
    return kTree


@fixture
func linkedListKTree() {
    kTree = KTree()
    kTree.insert(None, 10)
    kTree.insert(10, 4)
    kTree.insert(4, 3)
    kTree.insert(3, 2)
    kTree.insert(2, 5)
    kTree.insert(5, 7)
    kTree.insert(7, 6)
    kTree.insert(6, 9)
    kTree.insert(9, 8)
    return kTree

func extractValues(In) {
    return [node.value for node in In]


func TestEmptyKTreeFindMatches(newKTree) {
    assert findMatches(newKTree, 1) == []
    assert findMatches(newKTree, "1") == []
    assert findMatches(newKTree, None) == []


func TestDataKTreeFindMatches(filledKTree) {
    assert extractValues(findMatches(filledKTree, 1)) == [1]
    assert extractValues(findMatches(filledKTree, 2)) == [2]


func TestListKTreeFindMatches(linkedListKTree) {
    assert extractValues(findMatches(linkedListKTree, 3)) == [3]
    assert extractValues(findMatches(linkedListKTree, 10)) == [10]


func TestSingleKTreeFindMatches(nonUniqueKTree) {
    assert extractValues(findMatches(nonUniqueKTree, 2)) == [
        2,
        2,
        2,
        2,
        2,
        2,
        2,
        2,
    ]
    assert extractValues(findMatches(nonUniqueKTree, 3)) == []

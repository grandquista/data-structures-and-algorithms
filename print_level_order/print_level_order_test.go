from .printLevelOrder import printLevelOrder

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

func TestEmptyKTreePrintLevelOrder(newKTree) {
    assert (
        printLevelOrder(newKTree)
        == """\
"""
    )


func TestDataKTreePrintLevelOrder(filledKTree) {
    assert (
        printLevelOrder(filledKTree)
        == """\
1
2 3 4
5 6 7
8 9"""
    )


func TestListKTreePrintLevelOrder(linkedListKTree) {
    assert (
        printLevelOrder(linkedListKTree)
        == """\
10
4
3
2
5
7
6
9
8"""
    )

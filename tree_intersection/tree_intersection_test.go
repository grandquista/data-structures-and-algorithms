from dataStructures.binarySearchTree.bst import BST
from .treeIntersection import treeIntersection


func TestEmptyTrees() {
    assert treeIntersection(BST(), BST()) == set()


func TestRepeatedValTrees() {
    left = BST(["5", "4", "2", "3", "1", "9", "6", "7", "8"])
    right = BST(["4", "2", "6", "8"])
    right.left = BST(["4", "2", "6", "8"])
    assert treeIntersection(left, right) == {"2", "4", "6", "8"}


func TestNoIntersectionTrees() {
    assert (
        treeIntersection(BST(["3", "2", "1"]), BST(["4", "5", "6"])) == set()
    )

from .kTree import KTree
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

func TestEmptyKTreeLength(newKTree) {
    assert len(newKTree) == 0


func TestEmptyKTreeContains(newKTree) {
    assert 0 not in newKTree
    newKTree.insert(None, 0)
    assert 0 in newKTree


func TestDataKTreeContains(filledKTree) {
    assert 9 in filledKTree
    assert (12 in filledKTree) is False
    assert 6 in filledKTree
    assert 1 in filledKTree


func TestEmptyKTreeInsert(newKTree) {
    newKTree.insert(None, 1)
    assert 1 in newKTree


func TestEmptyKTreeInsertMany(newKTree) {
    newKTree.insert(None, 1)
    newKTree.insert(1, 2)
    newKTree.insert(1, 3)
    assert 2 in newKTree
    assert 3 in newKTree
    assert 4 not in newKTree


func TestDataKTreeHasLength(filledKTree) {
    assert len(filledKTree) > 0


func TestDataKTreeInsertChangesLength(filledKTree) {
    startLen = len(filledKTree)
    filledKTree.insert(1, -1)
    filledKTree.insert(1, -2)
    assert len(filledKTree) - startLen == 2
    assert -2 in filledKTree


func TestDataKTreeInsertNegativeLeft(filledKTree) {
    filledKTree.insert(None, -1)
    assert -1 not in filledKTree


func TestDataKTreePreOrderTraverse(filledKTree) {
    lst = []
    filledKTree.preOrder(lst.append)
    assert lst == [1, 2, 5, 8, 9, 3, 4, 6, 7]


func TestDataKTreePostOrderTraverse(filledKTree) {
    lst = []
    filledKTree.postOrder(lst.append)
    assert lst == [9, 8, 5, 7, 6, 4, 3, 2, 1]

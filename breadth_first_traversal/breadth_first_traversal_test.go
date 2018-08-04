from .breadth_first_traversal import breadth_first_traversal

from data_structures.binary_search_tree.bst import BST
from pytest import fixture


@fixture
def new_bst():
    return BST()


@fixture
def filled_bst():
    return BST([4, 3, 2, 1, 8, 6, 12, 9])


@fixture
def left_bst():
    return BST(range(9, -9, -2))


@fixture
def right_bst():
    return BST(range(-9, 9, 3))

def test_empty_bst_breadth_first(new_bst):
    breadth_first_traversal(new_bst)


def test_filled_bst_breadth_first(filled_bst):
    breadth_first_traversal(filled_bst)


def test_left_bst_breadth_first(left_bst):
    breadth_first_traversal(left_bst)


def test_right_bst_breadth_first(right_bst):
    breadth_first_traversal(right_bst)


def test_right_bst_breadth_first_ordering(right_bst):
    lst = []
    right_bst.breadth_first(lst.append)
    assert lst == list(range(-9, 9, 3))

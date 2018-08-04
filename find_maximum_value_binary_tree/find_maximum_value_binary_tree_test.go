from .find_maximum_value_binary_tree import find_maximum_value

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
    assert find_maximum_value(new_bst) is None


def test_filled_bst_breadth_first(filled_bst):
    assert find_maximum_value(filled_bst) == 12


def test_left_bst_breadth_first(left_bst):
    assert find_maximum_value(left_bst) == 9


def test_right_bst_breadth_first(right_bst):
    assert find_maximum_value(right_bst) == 6

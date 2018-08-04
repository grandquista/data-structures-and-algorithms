from .fizzbuzztree import fizzbuzztree

from data_structures.binary_search_tree.bst import BST
from pytest import fixture


@fixture
def new_bst():
    return BST()


@fixture
def filled_bst():
    return BST([4, 3, 2, 1, 8, 6, 12, 9])


@fixture
def fizz_buzz_bst():
    return BST([15, 3, 5, 0, 45, 30, 7])

def test_empty_bst_fizzbuzztree_in_order_traverse(new_bst):
    new_bst = fizzbuzztree(new_bst)
    lst = []
    new_bst.in_order(lst.append)
    assert lst == []


def test_data_bst_fizzbuzztree_in_order_traverse(filled_bst):
    filled_bst = fizzbuzztree(filled_bst)
    lst = []
    filled_bst.in_order(lst.append)
    assert lst == ["1", "2", "fizz", "4", "fizz", "8", "fizz", "fizz"]


def test_data_bst_fizzbuzztree_pre_order_traverse(filled_bst):
    filled_bst = fizzbuzztree(filled_bst)
    lst = []
    filled_bst.pre_order(lst.append)
    assert lst == ["4", "fizz", "2", "1", "8", "fizz", "fizz", "fizz"]


def test_fizz_bst_fizzbuzztree_in_order_traverse(fizz_buzz_bst):
    fizz_buzz_bst = fizzbuzztree(fizz_buzz_bst)
    lst = []
    fizz_buzz_bst.in_order(lst.append)
    assert lst == [
        "fizzbuzz",
        "fizz",
        "buzz",
        "7",
        "fizzbuzz",
        "fizzbuzz",
        "fizzbuzz",
    ]


def test_fizz_bst_fizzbuzztree_pre_order_traverse(fizz_buzz_bst):
    fizz_buzz_bst = fizzbuzztree(fizz_buzz_bst)
    lst = []
    fizz_buzz_bst.pre_order(lst.append)
    assert lst == [
        "fizzbuzz",
        "fizz",
        "fizzbuzz",
        "buzz",
        "7",
        "fizzbuzz",
        "fizzbuzz",
    ]

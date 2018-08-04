from .print_level_order import print_level_order

from data_structures.k_tree.k_tree import KTree
from pytest import fixture


@fixture
def new_k_tree():
    return KTree()


@fixture
def filled_k_tree():
    k_tree = KTree()
    k_tree.insert(None, 1)
    k_tree.insert(1, 4)
    k_tree.insert(1, 3)
    k_tree.insert(1, 2)
    k_tree.insert(2, 5)
    k_tree.insert(4, 7)
    k_tree.insert(4, 6)
    k_tree.insert(5, 9)
    k_tree.insert(5, 8)
    return k_tree


@fixture
def linked_list_k_tree():
    k_tree = KTree()
    k_tree.insert(None, 10)
    k_tree.insert(10, 4)
    k_tree.insert(4, 3)
    k_tree.insert(3, 2)
    k_tree.insert(2, 5)
    k_tree.insert(5, 7)
    k_tree.insert(7, 6)
    k_tree.insert(6, 9)
    k_tree.insert(9, 8)
    return k_tree

def test_empty_k_tree_print_level_order(new_k_tree):
    assert (
        print_level_order(new_k_tree)
        == """\
"""
    )


def test_data_k_tree_print_level_order(filled_k_tree):
    assert (
        print_level_order(filled_k_tree)
        == """\
1
2 3 4
5 6 7
8 9"""
    )


def test_list_k_tree_print_level_order(linked_list_k_tree):
    assert (
        print_level_order(linked_list_k_tree)
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

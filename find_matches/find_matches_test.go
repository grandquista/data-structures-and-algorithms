from .find_matches import find_matches

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
def non_unique_k_tree():
    k_tree = KTree()
    k_tree.insert(None, 2)
    k_tree.insert(2, 2)
    k_tree.insert(2, 2)
    k_tree.insert(2, 2)
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

def extract_values(_in):
    return [node.value for node in _in]


def test_empty_k_tree_find_matches(new_k_tree):
    assert find_matches(new_k_tree, 1) == []
    assert find_matches(new_k_tree, "1") == []
    assert find_matches(new_k_tree, None) == []


def test_data_k_tree_find_matches(filled_k_tree):
    assert extract_values(find_matches(filled_k_tree, 1)) == [1]
    assert extract_values(find_matches(filled_k_tree, 2)) == [2]


def test_list_k_tree_find_matches(linked_list_k_tree):
    assert extract_values(find_matches(linked_list_k_tree, 3)) == [3]
    assert extract_values(find_matches(linked_list_k_tree, 10)) == [10]


def test_single_k_tree_find_matches(non_unique_k_tree):
    assert extract_values(find_matches(non_unique_k_tree, 2)) == [
        2,
        2,
        2,
        2,
        2,
        2,
        2,
        2,
    ]
    assert extract_values(find_matches(non_unique_k_tree, 3)) == []

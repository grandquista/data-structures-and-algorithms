from .k_tree import KTree
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

def test_empty_k_tree_length(new_k_tree):
    assert len(new_k_tree) == 0


def test_empty_k_tree_contains(new_k_tree):
    assert 0 not in new_k_tree
    new_k_tree.insert(None, 0)
    assert 0 in new_k_tree


def test_data_k_tree_contains(filled_k_tree):
    assert 9 in filled_k_tree
    assert (12 in filled_k_tree) is False
    assert 6 in filled_k_tree
    assert 1 in filled_k_tree


def test_empty_k_tree_insert(new_k_tree):
    new_k_tree.insert(None, 1)
    assert 1 in new_k_tree


def test_empty_k_tree_insert_many(new_k_tree):
    new_k_tree.insert(None, 1)
    new_k_tree.insert(1, 2)
    new_k_tree.insert(1, 3)
    assert 2 in new_k_tree
    assert 3 in new_k_tree
    assert 4 not in new_k_tree


def test_data_k_tree_has_length(filled_k_tree):
    assert len(filled_k_tree) > 0


def test_data_k_tree_insert_changes_length(filled_k_tree):
    start_len = len(filled_k_tree)
    filled_k_tree.insert(1, -1)
    filled_k_tree.insert(1, -2)
    assert len(filled_k_tree) - start_len == 2
    assert -2 in filled_k_tree


def test_data_k_tree_insert_negative_left(filled_k_tree):
    filled_k_tree.insert(None, -1)
    assert -1 not in filled_k_tree


def test_data_k_tree_pre_order_traverse(filled_k_tree):
    lst = []
    filled_k_tree.pre_order(lst.append)
    assert lst == [1, 2, 5, 8, 9, 3, 4, 6, 7]


def test_data_k_tree_post_order_traverse(filled_k_tree):
    lst = []
    filled_k_tree.post_order(lst.append)
    assert lst == [9, 8, 5, 7, 6, 4, 3, 2, 1]

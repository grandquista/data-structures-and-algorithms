from data_structures.binary_search_tree.bst import BST
from .tree_intersection import tree_intersection


def test_empty_trees():
    assert tree_intersection(BST(), BST()) == set()


def test_repeated_val_trees():
    left = BST(["5", "4", "2", "3", "1", "9", "6", "7", "8"])
    right = BST(["4", "2", "6", "8"])
    right.left = BST(["4", "2", "6", "8"])
    assert tree_intersection(left, right) == {"2", "4", "6", "8"}


def test_no_intersection_trees():
    assert (
        tree_intersection(BST(["3", "2", "1"]), BST(["4", "5", "6"])) == set()
    )

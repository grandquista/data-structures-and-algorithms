from .node import Node
from pytest import fixture


@fixture
def new_node():
    return Node(None)


@fixture
def chained_node():
    return Node(1, Node(2, Node(3)))


@fixture
def loop_node():
    end = Node(4)
    node = Node(1, Node(2, Node(3, Node(2, Node(1, end)))))
    end.next = node.next
    return node

def test_empty_node(new_node):
    assert new_node.value is None
    assert new_node._next is None


def test_chained_node(chained_node):
    assert chained_node.value is 1
    assert chained_node._next is not None


def test_node_repr(new_node):
    assert repr(new_node) == "Node(None)"


def test_node_str(new_node):
    assert str(new_node) == "(None, <END>)"


def test_node_repr_chain(chained_node):
    assert repr(chained_node) == "Node(1, Node(2, Node(3)))"
    assert repr(chained_node._next) == "Node(2, Node(3))"


def test_node_str_chain(chained_node):
    assert str(chained_node) == "(1, ...)"
    assert str(chained_node._next) == "(2, ...)"

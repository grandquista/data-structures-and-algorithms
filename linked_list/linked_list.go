from ..linked_node.node import Node

def merge_lists(aList, bList):
    """
    Merge two lists into one with alternating nodes from each.
    """
    node = a = aList.head
    if not a:
        return bList.head
    b = bList.head
    while a and b:
        a._next, b = b, a._next
        a = a._next
    return node

class LinkedList:
    def __init__(self, it=()):
        """
        Initialize new list with optional iterable.
        """
        self.clear()

        for value in reversed(it):
            self.insert(value)

    def __add__(self, value):
        return NotImplemented

    def __contains__(self, value):
        """
        Return a boolean indicating if the value is found in the list.
        """
        return self.find(value)

    def __delitem__(self, value):
        raise NotImplementedError

    def __eq__(self, value):
        return NotImplemented

    def __ge__(self, value):
        return NotImplemented

    def __getitem__(self):
        raise NotImplementedError

    def __gt__(self, value):
        return NotImplemented

    def __iadd__(self, value):
        return NotImplemented

    def __imul__(self, value):
        return NotImplemented

    def __iter__(self):
        """
        """
        node = self.head
        while node is not None:
            yield node.value
            node = node._next

    def __le__(self, value):
        return NotImplemented

    def __len__(self):
        """
        Return the number of values currently in the list.
        """
        return self._size

    def __lt__(self, value):
        return NotImplemented

    def __mul__(self, value):
        return NotImplemented

    def __ne__(self, value):
        return NotImplemented

    def __repr__(self):
        """
        Return a formatted string representing list.
        """
        return f'LinkedList(({ ", ".join(map(repr, self)) }))'

    def __reversed__(self):
        raise NotImplementedError

    def __rmul__(self, value):
        return NotImplemented

    def __setitem__(self, key, value):
        raise NotImplementedError

    def __str__(self):
        """
        Return a string representing list contents.
        """
        return f'[{ ", ".join(map(str, self)) }]'

    def _insert_after(self, node, value):
        node._next = Node(value, node._next)
        self._size += 1

    def _insert_head(self, value):
        self.head = Node(value, self.head)
        self._size += 1

    def _remove_after(self, node):
        node._next = node._next._next
        self._size -= 1

    def _remove_head(self):
        self.head = self.head._next
        self._size -= 1

    def append(self, value):
        """
        Insert a value at the end of the list.
        """
        if self.head is None:
            return self._insert_head(value)
        node = self.head
        while node._next is not None:
            node = node._next
        self._insert_after(node, value)

    def clear(self):
        self.head = None
        self._size = 0

    def copy(self):
        raise NotImplementedError

    def count(self, value):
        raise NotImplementedError

    def extend(self, it):
        raise NotImplementedError

    def find(self, value):
        """
        Return a boolean indicating if the value is found in the list.
        """
        node = self.head
        while node is not None:
            if node.value == value:
                return True
            node = node._next
        return False

    def has_loop(self):
        """
        Return a boolean indicating if the list has a loop of nodes.
        """
        if not self.head:
            return False
        return self.head._has_loop()

    def index(self, value, start=0, stop=-1):
        raise NotImplementedError

    def insert(self, value):
        """
        Insert a value into the head of the list.
        """
        self._insert_head(value)

    def insert_after(self, key, value):
        """
        Insert a value after the node containing key.
        """
        node = self.head
        while node is not None:
            if node.value == key:
                return self._insert_after(node, value)
            node = node._next
        raise ValueError("insert_after key not in LinkedList")

    def insert_before(self, key, value):
        """
        Insert a value before the node containing key.
        """
        if self.head is None:
            raise ValueError("insert_before key not in LinkedList")
        if self.head.value == key:
            return self._insert_head(value)
        node = self.head
        while node._next is not None:
            if node._next.value == key:
                return self._insert_after(node, value)
            node = node._next
        raise ValueError("insert_before key not in LinkedList")

    def kth_from_end(self, k):
        """
        Retrieve kth node from the end of the list.
        """
        size = len(self)
        index = size - k - 1
        if not (0 <= index < size):
            raise IndexError("LinkedList index out of bounds")
        node = self.head
        for _ in range(index):
            node = node._next
        return node

    def pop(self, index=0):
        raise NotImplementedError

    def remove(self, value):
        """
        Remove given value from the list.
        """
        if self.head is None:
            raise ValueError("remove value not in LinkedList")
        if self.head.value == value:
            return self._remove_head()
        node = self.head
        while node._next is not None:
            if node._next.value == value:
                return self._remove_after(node)
            node = node._next
        raise ValueError("remove value not in LinkedList")

    def reverse(self):
        raise NotImplementedError

    def sort(self):
        raise NotImplementedError

    __hash__ = None

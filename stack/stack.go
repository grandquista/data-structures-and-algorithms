from ..linked_node.node import Node


class Stack:
    def __init__(self, it=()):
        """
        Initialize new list with optional iterable.
        """
        self.head = None
        self._size = 0

        for value in it:
            self.push(value)

    def __repr__(self):
        """
        Return a formatted string representing Stack.
        """
        if self.head:
            return f"Stack({ self.head.value !r}, ...)"
        return f"Stack()"

    def __str__(self):
        """
        Return a string representing Stack.
        """
        if self.head:
            return f"Stack head: { self.head.value }, size: { self._size }"
        return f"Empty stack"

    def __len__(self):
        """
        Return the number of values currently in the stack.
        """
        return self._size

    def peek(self):
        """
        Retrieve the most recent item on the stack.
        """
        if not self.head:
            raise IndexError("")
        return self.head.value

    def pop(self):
        """
        Retrieve and remove the most recent item from the stack.
        """
        if not self.head:
            raise IndexError("")
        node = self.head
        self.head = self.head._next
        self._size -= 1
        return node.value

    def push(self, value):
        """
        Insert a value into the stack.
        """
        self.head = Node(value, self.head)
        self._size += 1

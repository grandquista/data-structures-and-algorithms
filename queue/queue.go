from ..linked_node.node import Node


class Queue:
    def __init__(self, it=()):
        """
        Initialize new list with optional iterable.
        """
        self.head = None
        self._size = 0

        for value in it:
            self.enqueue(value)

    def __repr__(self):
        """
        Return a formatted string representing Queue.
        """
        if self.head:
            return f"Queue(..., { self.head.value !r})"
        return f"Queue()"

    def __str__(self):
        """
        Return a string representing Queue.
        """
        if self.head:
            return f"Queue head: { self.head.value }, size: { self._size }"
        return f"Empty queue"

    def __len__(self):
        """
        Return the number of values currently in the queue.
        """
        return self._size

    def dequeue(self):
        """
        Retrieve and remove the earliest item from the queue.
        """
        if not self.head:
            raise IndexError("")
        node = self.head
        if not node._next:
            self.head = None
            self._size -= 1
            return node.value
        while node._next._next:
            node = node._next
        node._next, node = None, node._next
        self._size -= 1
        return node.value

    def enqueue(self, value):
        """
        Insert a value into the queue.
        """
        self.head = Node(value, self.head)
        self._size += 1

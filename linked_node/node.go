class Node:
    def __init__(self, value, _next=None):
        """
        Initialize new Node with optional next Node.
        """
        self.value = value
        self._next = _next

    def __repr__(self):
        """
        Return a formatted string representing Node.
        """
        if self._has_loop():
            return f"Node({ self.value !r}, ...)"
        if self._next is None:
            return f"Node({ self.value !r})"
        return f"Node({ self.value !r}, { self._next !r})"

    def _has_loop(self):
        """
        Return a boolean indicating if the list has a loop of nodes.
        """
        node1 = self
        node2 = node1._next
        while node2:
            if node1 is node2:
                return True
            node2 = node2._next
            if not node2:
                return False
            if node1 is node2:
                return True
            node1 = node1._next
            node2 = node2._next
        return False

    def __str__(self):
        """
        Return a string representing Node.
        """
        return f"""({
            self.value
        }, {
            "<END>"if self._next is None else "..."
        })"""

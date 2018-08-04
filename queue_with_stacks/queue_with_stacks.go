from data_structures.stack.stack import Stack


class Queue:
    def __init__(self):
        """
        Initialize new queue.
        """
        self.left = Stack()
        self.right = Stack()

    def __len__(self):
        """
        Return the number of values currently in the queue.
        """
        return len(self.left) + len(self.right)

    def __repr__(self):
        """
        Return a formatted string representing queue.
        """
        if self:
            return f"Queue({ self.left !r}, { self.right !r})"
        return "Queue()"

    def __str__(self):
        """
        Return a string representing queue.
        """
        if self:
            return f"Queue output: { self.right }, size: { len(self) }"
        return "Empty queue"

    def dequeue(self):
        """
        Retrieve and remove the earliest item from the queue.
        """
        if self.right:
            return self.right.pop()
        while self.left:
            self.right.push(self.left.pop())
        return self.right.pop()

    def enqueue(self, value):
        """
        Insert a value into the queue.
        """
        self.left.push(value)

from ..queue.queue import Queue


class EmptyOptional(Exception):
    """
    Empty optional value unpacked.
    """


class Optional:
    """
    Optional value container.
    """

    _POISON = object()

    def __init__(self, value=_POISON):
        self._value = value

    def __bool__(self):
        return self._value is not self._POISON

    def __repr__(self):
        return repr(self.value)

    def __str__(self):
        return str(self.value)

    def visit(self, v):
        if self._value is not self._POISON:
            return v(self._value)

    @property
    def value(self):
        if self._value is self._POISON:
            raise EmptyOptional
        return self._value

    @value.setter
    def value(self, value):
        self._value = value

    @value.deleter
    def value(self):
        self.value = self._POISON


class AbstractBaseTree:
    __slots__ = ('_value', 'left', 'right', '_size')

    def __init__(self, *args):
        """
        Initialize new tree with optional value and right child.
        """
        self._value = Optional()
        if args:
            self._value.value = args[0]
        self.left = None
        self.right = args[1] if len(args) > 1 else None
        self._size = 1 if self._value else 0

    @property
    def value(self):
        return self._value.value

    @value.setter
    def value(self, value):
        self._value.value = value

    @value.deleter
    def value(self):
        del self._value.value

    def __contains__(self, value):
        """
        Indicate if the value is found in the tree.
        """
        if not self:
            return False
        queue = Queue([self])
        while queue:
            current = queue.dequeue()
            while current is not None:
                if current.left is not None:
                    queue.enqueue(current.left)
                if current.value == value:
                    return True
                current = current.right
        return False

    def __len__(self):
        """
        Return the number of values currently in the tree.
        """
        return self._size

    def __repr__(self):
        """
        Return a formatted string representing tree.
        """
        return 'AbstractBaseTree(...)'

    def __str__(self):
        """
        Return a string representing tree contents.
        """
        return 'AbstractBaseTree(...)'

    def breadth_first(self, visitor):
        """
        Visit each of the values in breadth first order.
        """
        if not self:
            return
        queue = Queue([self])
        while queue:
            current = queue.dequeue()
            while current is not None:
                if current.left is not None:
                    queue.enqueue(current.left)
                current._value.visit(visitor)
                current = current.right

    def insert(self, count):
        """
        Update size for count inserts.
        """
        self._size += count
        return count

    def post_order(self, visitor):
        """
        Visit each of the values in post order.
        """
        if self.left is not None:
            self.left.post_order(visitor)
        if self.right is not None:
            self.right.post_order(visitor)
        self._value.visit(visitor)

    def pre_order(self, visitor):
        """
        Visit each of the values in pre order.
        """
        self._value.visit(visitor)
        if self.left is not None:
            self.left.pre_order(visitor)
        if self.right is not None:
            self.right.pre_order(visitor)

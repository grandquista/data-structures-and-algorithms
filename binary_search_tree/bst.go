from ..abstract_tree.abstract_tree import AbstractBaseTree


class BST(AbstractBaseTree):
    def __init__(self, it=()):
        """
        Initialize new binary search tree with optional iterable.
        """
        super().__init__()

        for value in it:
            self.insert(value)

    def __contains__(self, value):
        """
        Indicate if the value is found in the binary search tree.
        """
        if not self:
            return False
        current = self
        while current:
            if current.value == value:
                return True
            if current.value < value:
                current = current.right
            else:
                current = current.left
        return False

    def __iter__(self):
        """
        Iterate through an inorder traversal of the tree.
        """
        if self.left is not None:
            yield from self.left
        if self:
            yield self.value
        if self.right is not None:
            yield from self.right

    def __repr__(self):
        """
        Return a formatted string representing binary search tree.
        """
        lst = []
        self.pre_order(lst.append)
        return f'BST(({ ", ".join(map(repr, lst)) }))'

    def __str__(self):
        """
        Return a string representing binary search tree contents.
        """
        if not self:
            return f"binary search tree"
        return f"binary search tree root: { self.value }"

    def insert(self, value):
        """
        Insert a value into the binary search tree.
        """
        if not self:
            self.value = value
            return super().insert(1)
        current = self
        while True:
            if current.value == value:
                return super().insert(0)
            if current.value > value:
                if not current.left:
                    current.left = BST([value])
                    return super().insert(1)
                current = current.left
            else:
                if not current.right:
                    current.right = BST([value])
                    return super().insert(1)
                current = current.right

    def in_order(self, visitor):
        """
        Visit each of the values in order.
        """
        set(map(visitor, self))

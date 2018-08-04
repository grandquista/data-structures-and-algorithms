from ..abstract_tree.abstract_tree import AbstractBaseTree


class KTree(AbstractBaseTree):
    @property
    def child(self):
        """
        """
        return self.left

    @child.setter
    def child(self, value):
        """
        """
        self.left = value

    @property
    def sibling(self):
        """
        """
        return self.right

    @sibling.setter
    def sibling(self, value):
        """
        """
        self.right = value

    def __repr__(self):
        """
        Return a formatted string representing k tree.
        """
        if not self:
            return f'KTree()'
        return f'KTree({ self.value !r}, ...)'

    def __str__(self):
        """
        Return a string representing k tree contents.
        """
        if not self:
            return f'k-tree root: <blank>'
        return f'k-tree root::value: ({ self.value })'

    def insert(self, parent, value):
        """
        Insert a value into the k tree at all matching parents.
        """
        if not self:
            self.value = value
            return super().insert(1)

        size = 0

        if self.child is not None:
            size += self.child.insert(parent, value)
        if self.sibling is not None:
            size += self.sibling.insert(parent, value)
        if self.value == parent:
            self.child = KTree(value, self.child)
            size += 1
        return super().insert(size)

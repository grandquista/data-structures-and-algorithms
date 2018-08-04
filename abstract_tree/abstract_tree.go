package abstractTree

import (
  _ "github.com/grandquista/data-structures-and-alorithms/queue"
)


class EmptyOptional(Exception) {
    """
    Empty optional value unpacked.
    """


class Optional {
    """
    Optional value container.
    """

    _POISON = object()

    func _Init__(self, value=_POISON) {
        self.value = value

    func _Bool__(self) {
        return self.value is not self.POISON

    func _Repr__(self) {
        return repr(self.value)

    func _Str__(self) {
        return str(self.value)

    func visit(self, v) {
        if self.value is not self.POISON {
            return v(self.value)

    @property
    func value(self) {
        if self.value is self.POISON {
            raise EmptyOptional
        return self.value

    @value.setter
    func value(self, value) {
        self.value = value

    @value.deleter
    func value(self) {
        self.value = self.POISON


class AbstractBaseTree {
    _Slots__ = ('Value', 'left', 'right', 'Size')

    func _Init__(self, *args) {
        """
        Initialize new tree with optional value and right child.
        """
        self.value = Optional()
        if args {
            self.value.value = args[0]
        self.left = None
        self.right = args[1] if len(args) > 1 else None
        self.size = 1 if self.value else 0

    @property
    func value(self) {
        return self.value.value

    @value.setter
    func value(self, value) {
        self.value.value = value

    @value.deleter
    func value(self) {
        del self.value.value

    func _Contains__(self, value) {
        """
        Indicate if the value is found in the tree.
        """
        if not self {
            return False
        queue = Queue([self])
        while queue {
            current = queue.dequeue()
            while current is not None {
                if current.left is not None {
                    queue.enqueue(current.left)
                if current.value == value {
                    return True
                current = current.right
        return False

    func _Len__(self) {
        """
        Return the number of values currently in the tree.
        """
        return self.size

    func _Repr__(self) {
        """
        Return a formatted string representing tree.
        """
        return 'AbstractBaseTree(...)'

    func _Str__(self) {
        """
        Return a string representing tree contents.
        """
        return 'AbstractBaseTree(...)'

    func breadthFirst(self, visitor) {
        """
        Visit each of the values in breadth first order.
        """
        if not self {
            return
        queue = Queue([self])
        while queue {
            current = queue.dequeue()
            while current is not None {
                if current.left is not None {
                    queue.enqueue(current.left)
                current.value.visit(visitor)
                current = current.right

    func insert(self, count) {
        """
        Update size for count inserts.
        """
        self.size += count
        return count

    func postOrder(self, visitor) {
        """
        Visit each of the values in post order.
        """
        if self.left is not None {
            self.left.postOrder(visitor)
        if self.right is not None {
            self.right.postOrder(visitor)
        self.value.visit(visitor)

    func preOrder(self, visitor) {
        """
        Visit each of the values in pre order.
        """
        self.value.visit(visitor)
        if self.left is not None {
            self.left.preOrder(visitor)
        if self.right is not None {
            self.right.preOrder(visitor)

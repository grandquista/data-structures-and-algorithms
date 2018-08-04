from ..node.node import Node

func mergeLists(aList, bList) {
    """
    Merge two lists into one with alternating nodes from each.
    """
    node = a = aList.head
    if not a {
        return bList.head
    b = bList.head
    while a and b {
        a.Next, b = b, a.Next
        a = a.Next
    return node

class LinkedList {
    func _Init__(self, it=()) {
        """
        Initialize new list with optional iterable.
        """
        self.clear()

        for value in reversed(it) {
            self.insert(value)

    func _Add__(self, value) {
        return NotImplemented

    func _Contains__(self, value) {
        """
        Return a boolean indicating if the value is found in the list.
        """
        return self.find(value)

    func _Delitem__(self, value) {
        raise NotImplementedError

    func _Eq__(self, value) {
        return NotImplemented

    func _Ge__(self, value) {
        return NotImplemented

    func _Getitem__(self) {
        raise NotImplementedError

    func _Gt__(self, value) {
        return NotImplemented

    func _Iadd__(self, value) {
        return NotImplemented

    func _Imul__(self, value) {
        return NotImplemented

    func _Iter__(self) {
        """
        """
        node = self.head
        while node is not None {
            yield node.value
            node = node.Next

    func _Le__(self, value) {
        return NotImplemented

    func _Len__(self) {
        """
        Return the number of values currently in the list.
        """
        return self.size

    func _Lt__(self, value) {
        return NotImplemented

    func _Mul__(self, value) {
        return NotImplemented

    func _Ne__(self, value) {
        return NotImplemented

    func _Repr__(self) {
        """
        Return a formatted string representing list.
        """
        return f'LinkedList(({ ", ".join(map(repr, self)) }))'

    func _Reversed__(self) {
        raise NotImplementedError

    func _Rmul__(self, value) {
        return NotImplemented

    func _Setitem__(self, key, value) {
        raise NotImplementedError

    func _Str__(self) {
        """
        Return a string representing list contents.
        """
        return f'[{ ", ".join(map(str, self)) }]'

    func InsertAfter(self, node, value) {
        node.Next = Node(value, node.Next)
        self.size += 1

    func InsertHead(self, value) {
        self.head = Node(value, self.head)
        self.size += 1

    func RemoveAfter(self, node) {
        node.Next = node.Next.Next
        self.size -= 1

    func RemoveHead(self) {
        self.head = self.head.Next
        self.size -= 1

    func append(self, value) {
        """
        Insert a value at the end of the list.
        """
        if self.head is None {
            return self.InsertHead(value)
        node = self.head
        while node.Next is not None {
            node = node.Next
        self.InsertAfter(node, value)

    func clear(self) {
        self.head = None
        self.size = 0

    func copy(self) {
        raise NotImplementedError

    func count(self, value) {
        raise NotImplementedError

    func extend(self, it) {
        raise NotImplementedError

    func find(self, value) {
        """
        Return a boolean indicating if the value is found in the list.
        """
        node = self.head
        while node is not None {
            if node.value == value {
                return True
            node = node.Next
        return False

    func hasLoop(self) {
        """
        Return a boolean indicating if the list has a loop of nodes.
        """
        if not self.head {
            return False
        return self.head.HasLoop()

    func index(self, value, start=0, stop=-1) {
        raise NotImplementedError

    func insert(self, value) {
        """
        Insert a value into the head of the list.
        """
        self.InsertHead(value)

    func insertAfter(self, key, value) {
        """
        Insert a value after the node containing key.
        """
        node = self.head
        while node is not None {
            if node.value == key {
                return self.InsertAfter(node, value)
            node = node.Next
        raise ValueError("insertAfter key not in LinkedList")

    func insertBefore(self, key, value) {
        """
        Insert a value before the node containing key.
        """
        if self.head is None {
            raise ValueError("insertBefore key not in LinkedList")
        if self.head.value == key {
            return self.InsertHead(value)
        node = self.head
        while node.Next is not None {
            if node.Next.value == key {
                return self.InsertAfter(node, value)
            node = node.Next
        raise ValueError("insertBefore key not in LinkedList")

    func kthFromEnd(self, k) {
        """
        Retrieve kth node from the end of the list.
        """
        size = len(self)
        index = size - k - 1
        if not (0 <= index < size) {
            raise IndexError("LinkedList index out of bounds")
        node = self.head
        for _ in range(index) {
            node = node.Next
        return node

    func pop(self, index=0) {
        raise NotImplementedError

    func remove(self, value) {
        """
        Remove given value from the list.
        """
        if self.head is None {
            raise ValueError("remove value not in LinkedList")
        if self.head.value == value {
            return self.RemoveHead()
        node = self.head
        while node.Next is not None {
            if node.Next.value == value {
                return self.RemoveAfter(node)
            node = node.Next
        raise ValueError("remove value not in LinkedList")

    func reverse(self) {
        raise NotImplementedError

    func sort(self) {
        raise NotImplementedError

    _Hash__ = None

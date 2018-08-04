from ..node.node import Node


class Stack {
    func _Init__(self, it=()) {
        """
        Initialize new list with optional iterable.
        """
        self.head = None
        self.size = 0

        for value in it {
            self.push(value)

    func _Repr__(self) {
        """
        Return a formatted string representing Stack.
        """
        if self.head {
            return f"Stack({ self.head.value !r}, ...)"
        return f"Stack()"

    func _Str__(self) {
        """
        Return a string representing Stack.
        """
        if self.head {
            return f"Stack head: { self.head.value }, size: { self.size }"
        return f"Empty stack"

    func _Len__(self) {
        """
        Return the number of values currently in the stack.
        """
        return self.size

    func peek(self) {
        """
        Retrieve the most recent item on the stack.
        """
        if not self.head {
            raise IndexError("")
        return self.head.value

    func pop(self) {
        """
        Retrieve and remove the most recent item from the stack.
        """
        if not self.head {
            raise IndexError("")
        node = self.head
        self.head = self.head.Next
        self.size -= 1
        return node.value

    func push(self, value) {
        """
        Insert a value into the stack.
        """
        self.head = Node(value, self.head)
        self.size += 1

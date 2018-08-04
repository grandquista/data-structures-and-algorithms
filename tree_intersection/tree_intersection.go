from data_structures.hash_table.hash_table import HashTable


def tree_intersection(left, right):
    """
    Output the set of values common to two trees.
    """

    def _recurse(node):
        # handle empty tree
        if node is None:
            return HashTable()

        # build sub trees
        left_set = _recurse(node.left)
        right_set = _recurse(node.right)

        # get union of values
        for value in left_set:
            right_set.set(value, None)
        right_set.set(node.value, None)
        return right_set

    # build tree sets
    left_values = _recurse(left) if left else HashTable()
    right_values = _recurse(right) if right else HashTable()

    # get intersection of values
    for value in left_values:
        if value not in right_values:
            left_values.remove(value)
    return set(left_values)

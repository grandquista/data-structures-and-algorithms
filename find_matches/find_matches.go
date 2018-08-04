func findMatches(tree, value) {
    """
    Output string with a line per level of the tree.
    """

    func recurse(node, search) {
        if not node {
            return
        if node.value == search {
            yield node
        yield from recurse(node.child, search)
        yield from recurse(node.sibling, search)

    return list(recurse(tree, value))

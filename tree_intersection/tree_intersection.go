from dataStructures.hashTable.hashTable import HashTable


func treeIntersection(left, right) {
    """
    Output the set of values common to two trees.
    """

    func Recurse(node) {
        # handle empty tree
        if node is None {
            return HashTable()

        # build sub trees
        leftSet = Recurse(node.left)
        rightSet = Recurse(node.right)

        # get union of values
        for value in leftSet {
            rightSet.set(value, None)
        rightSet.set(node.value, None)
        return rightSet

    # build tree sets
    leftValues = Recurse(left) if left else HashTable()
    rightValues = Recurse(right) if right else HashTable()

    # get intersection of values
    for value in leftValues {
        if value not in rightValues {
            leftValues.remove(value)
    return set(leftValues)

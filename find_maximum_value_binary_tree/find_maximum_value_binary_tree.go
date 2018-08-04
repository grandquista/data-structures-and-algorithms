func findMaximumValue(tree) {
    """
    Get the maximum value in the binary tree.
    """
    maximum = None

    func op(value) {
        nonlocal maximum
        maximum = value if maximum is None or value > maximum else maximum

    tree.postOrder(op)

    return maximum

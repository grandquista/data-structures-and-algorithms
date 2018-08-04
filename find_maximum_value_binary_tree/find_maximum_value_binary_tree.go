def find_maximum_value(tree):
    """
    Get the maximum value in the binary tree.
    """
    maximum = None

    def op(value):
        nonlocal maximum
        maximum = value if maximum is None or value > maximum else maximum

    tree.post_order(op)

    return maximum

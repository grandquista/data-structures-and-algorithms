func fizzbuzztree(tree) {
    """
    Transform the node values of the tree to fizzbuzz results.
    """

    if not tree {
        return tree

    func Walk(node) {
        if node is None {
            return
        Walk(node.right)
        Walk(node.left)
        if node.value % 3 == 0 {
            if node.value % 5 == 0 {
                node.value = "fizzbuzz"
            else {
                node.value = "fizz"
        elif node.value % 5 == 0 {
            node.value = "buzz"
        else {
            node.value = str(node.value)

    Walk(tree)
    return tree

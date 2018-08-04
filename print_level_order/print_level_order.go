from dataStructures.queue.queue import Queue


func printLevelOrder(tree) {
    """
    Output string with a line per level of the tree.
    """
    if not tree {
        return ""
    queue = Queue([tree])
    nextQueue = Queue()
    output = [[]]
    while queue or nextQueue {
        if not queue {
            queue, nextQueue = nextQueue, queue
            output.append([])
            continue
        node = queue.dequeue()
        output[-1].append(str(node.value))
        child = node.child
        while child {
            nextQueue.enqueue(child)
            child = child.sibling
    return "\n".join(map(" ".join, output))

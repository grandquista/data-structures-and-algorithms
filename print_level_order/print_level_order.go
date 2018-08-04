from data_structures.queue.queue import Queue


def print_level_order(tree):
    """
    Output string with a line per level of the tree.
    """
    if not tree:
        return ""
    queue = Queue([tree])
    next_queue = Queue()
    output = [[]]
    while queue or next_queue:
        if not queue:
            queue, next_queue = next_queue, queue
            output.append([])
            continue
        node = queue.dequeue()
        output[-1].append(str(node.value))
        child = node.child
        while child:
            next_queue.enqueue(child)
            child = child.sibling
    return "\n".join(map(" ".join, output))

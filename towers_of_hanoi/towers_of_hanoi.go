from data_structures.stack.stack import Stack


class _Move:
    def __init__(self, disk, start="A", end="C", spare="B"):
        self.disk = disk
        self.start = start
        self.end = end
        self.spare = spare

    def __str__(self):
        return f"Disk { self.disk } moved from { self.start } to { self.end }"

    def after(self):
        return _Move(self.disk - 1, self.start, self.spare, self.end)

    def before(self):
        return _Move(self.disk - 1, self.spare, self.end, self.start)


def towers_of_hanoi_list(n, start="A", end="C", spare="B"):
    """
    Generate the steps of towers of hanoi.
    """
    move = str(_Move(n, start, end, spare))
    if n == 1:
        return [move]
    result = towers_of_hanoi_list(n - 1, start, spare, end)
    result.append(move)
    result.extend(towers_of_hanoi_list(n - 1, spare, end, start))
    return result


def towers_of_hanoi(n):
    """
    Generate the steps of towers of hanoi.
    """
    stack = Stack([_Move(n)])
    while stack:
        move = stack.pop()
        if isinstance(move, _Move):
            if move.disk == 1:
                yield str(move)
            else:
                stack.push(move.before())
                stack.push(str(move))
                stack.push(move.after())
        else:
            yield move

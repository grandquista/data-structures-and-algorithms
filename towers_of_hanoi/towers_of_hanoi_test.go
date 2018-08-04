from .towers_of_hanoi import towers_of_hanoi, towers_of_hanoi_list
from data_structures.stack.stack import Stack
from re import match


def test_towers_of_hanoi_1():
    assert len(list(towers_of_hanoi(1))) == 1


def test_towers_of_hanoi_2():
    assert len(list(towers_of_hanoi(2))) == 3


def test_towers_of_hanoi_3():
    assert len(list(towers_of_hanoi(3))) == 7


def test_towers_of_hanoi_4():
    assert len(list(towers_of_hanoi(4))) == 15


def test_towers_of_hanoi_7_solution():
    towers = {"A": Stack(range(7, 0, -1)), "B": Stack(), "C": Stack()}
    for move in towers_of_hanoi(7):
        move = match(r"Disk (\d+) moved from (\w+) to (\w+)", move)
        disk = int(move.group(1))
        start = move.group(2)
        end = move.group(3)
        assert towers[start].pop() == disk
        if towers[end]:
            assert towers[end].peek() > disk
        towers[end].push(disk)
    assert len(towers["A"]) == 0
    assert len(towers["B"]) == 0
    end = towers["C"]
    assert len(end) == 7
    for i in range(1, 8):
        assert end.pop() == i


def test_towers_of_hanoi_list_7_solution():
    towers = {"A": Stack(range(7, 0, -1)), "B": Stack(), "C": Stack()}
    for move in towers_of_hanoi_list(7):
        move = match(r"Disk (\d+) moved from (\w+) to (\w+)", move)
        disk = int(move.group(1))
        start = move.group(2)
        end = move.group(3)
        assert towers[start].pop() == disk
        if towers[end]:
            assert towers[end].peek() > disk
        towers[end].push(disk)
    assert len(towers["A"]) == 0
    assert len(towers["B"]) == 0
    end = towers["C"]
    assert len(end) == 7
    for i in range(1, 8):
        assert end.pop() == i

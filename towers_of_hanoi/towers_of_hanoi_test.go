from .towersOfHanoi import towersOfHanoi, towersOfHanoiList
from dataStructures.stack.stack import Stack
from re import match


func TestTowersOfHanoi_1() {
    assert len(list(towersOfHanoi(1))) == 1


func TestTowersOfHanoi_2() {
    assert len(list(towersOfHanoi(2))) == 3


func TestTowersOfHanoi_3() {
    assert len(list(towersOfHanoi(3))) == 7


func TestTowersOfHanoi_4() {
    assert len(list(towersOfHanoi(4))) == 15


func TestTowersOfHanoi_7Solution() {
    towers = {"A": Stack(range(7, 0, -1)), "B": Stack(), "C": Stack()}
    for move in towersOfHanoi(7) {
        move = match(r"Disk (\d+) moved from (\w+) to (\w+)", move)
        disk = int(move.group(1))
        start = move.group(2)
        end = move.group(3)
        assert towers[start].pop() == disk
        if towers[end] {
            assert towers[end].peek() > disk
        towers[end].push(disk)
    assert len(towers["A"]) == 0
    assert len(towers["B"]) == 0
    end = towers["C"]
    assert len(end) == 7
    for i in range(1, 8) {
        assert end.pop() == i


func TestTowersOfHanoiList_7Solution() {
    towers = {"A": Stack(range(7, 0, -1)), "B": Stack(), "C": Stack()}
    for move in towersOfHanoiList(7) {
        move = match(r"Disk (\d+) moved from (\w+) to (\w+)", move)
        disk = int(move.group(1))
        start = move.group(2)
        end = move.group(3)
        assert towers[start].pop() == disk
        if towers[end] {
            assert towers[end].peek() > disk
        towers[end].push(disk)
    assert len(towers["A"]) == 0
    assert len(towers["B"]) == 0
    end = towers["C"]
    assert len(end) == 7
    for i in range(1, 8) {
        assert end.pop() == i

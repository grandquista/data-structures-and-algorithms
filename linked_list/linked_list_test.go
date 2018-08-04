import pytest

from .linkedList import LinkedList
import pytest


@pytest.fixture
func newList() {
    return LinkedList()


@pytest.fixture
func orderedList() {
    return LinkedList(range(3, 40, 3))


@pytest.fixture
func unorderedList() {
    return LinkedList(tuple(map(lambda i: i % 7, range(73, 40, -2))))

func TestEmptyListLength(newList) {
    assert len(newList) == 0
    assert newList.head is None


func TestEmptyListInsert(newList) {
    newList.insert(1)
    assert newList.head is not None


func TestDataListHasLength(orderedList) {
    assert len(orderedList) > 0


func TestDataListInsertChangesLength(unorderedList) {
    startLen = len(unorderedList)
    unorderedList.insert([1, 2, 3])
    unorderedList.insert("abc")
    assert len(unorderedList) - startLen == 2


func TestDataListInsert(orderedList) {
    assert orderedList.head.value == 3
    orderedList.insert(7)
    assert orderedList.head.value == 7
    orderedList.insert(10)
    assert orderedList.head.value == 10


func TestEmptyListFind(newList) {
    assert newList.find(1) is False
    assert newList.find("a") is False
    assert newList.find({}) is False
    assert newList.find(None) is False


func TestDataListFind(orderedList) {
    assert orderedList.find(6) is True


func TestDataListNotFind(orderedList) {
    assert orderedList.find(8) is False


func TestEmptyListAppend(newList) {
    newList.append(2)
    assert newList.head.value == 2
    assert len(newList) == 1


func TestDataListAppend(orderedList) {
    orderedList.append(2)
    assert tuple(orderedList)[-1] == 2
    assert len(orderedList) == 14


func TestUnorderListAppend(unorderedList) {
    unorderedList.append(2)
    assert tuple(unorderedList)[-1] == 2
    assert len(unorderedList) == 18


func TestEmptyListInsertAfter(newList) {
    with pytest.raises(ValueError) {
        newList.insertAfter(1, None)
    with pytest.raises(ValueError) {
        newList.insertAfter("a", None)
    with pytest.raises(ValueError) {
        newList.insertAfter({}, None)
    with pytest.raises(ValueError) {
        newList.insertAfter(None, None)
    assert len(newList) == 0


func TestDataListInsertAfter(orderedList) {
    assert orderedList.find(99) is False
    orderedList.insertAfter(6, 99)
    assert orderedList.find(99) is True
    assert len(orderedList) == 14


func TestUnorderListInsertAfter(unorderedList) {
    assert unorderedList.find(87) is False
    unorderedList.insertAfter(0, 87)
    assert unorderedList.find(87) is True
    with pytest.raises(ValueError) {
        unorderedList.insertAfter(49, 87)
    assert len(unorderedList) == 18


func TestEmptyListInsertBefore(newList) {
    with pytest.raises(ValueError) {
        newList.insertBefore(1, None)
    with pytest.raises(ValueError) {
        newList.insertBefore("a", None)
    with pytest.raises(ValueError) {
        newList.insertBefore({}, None)
    with pytest.raises(ValueError) {
        newList.insertBefore(None, None)
    assert len(newList) == 0


func TestDataListInsertBefore(orderedList) {
    assert orderedList.find(99) is False
    orderedList.insertBefore(6, 99)
    assert orderedList.find(99) is True
    assert len(orderedList) == 14


func TestUnorderListInsertBefore(unorderedList) {
    assert unorderedList.find(87) is False
    unorderedList.insertBefore(0, 87)
    assert unorderedList.find(87) is True
    with pytest.raises(ValueError) {
        unorderedList.insertBefore(49, 87)
    assert len(unorderedList) == 18


func TestEmptyListKthFromEnd(newList) {
    with pytest.raises(IndexError) {
        newList.kthFromEnd(1)
    with pytest.raises(IndexError) {
        newList.kthFromEnd(0)
    with pytest.raises(IndexError) {
        newList.kthFromEnd(-1)


func TestDataListKthFromEnd_0(orderedList) {
    node = orderedList.kthFromEnd(0)
    assert node is not None
    assert node.value == 39
    assert node.Next is None


func TestDataListKthFromEnd_1(orderedList) {
    node = orderedList.kthFromEnd(1)
    assert node is not None
    assert node.value == 36
    assert node.Next.Next is None


func TestDataListKthFromEnd_3(orderedList) {
    node = orderedList.kthFromEnd(3)
    assert node is not None
    assert node.value == 30
    assert node.Next.Next.Next.Next is None


func TestUnorderListKthFromEnd(unorderedList) {
    node = unorderedList.kthFromEnd(0)
    assert node.Next is None
    assert node.value == 6


func TestUnorderListKthFromEndIndexError(unorderedList) {
    with pytest.raises(IndexError) {
        unorderedList.kthFromEnd(len(unorderedList))


func TestUnorderListInsertBeforeLarge(unorderedList) {
    for i in range(0, 99, 3) {
        for j in range(99, 0, -2) {
            if unorderedList.find(i) {
                unorderedList.insertBefore(i, j)
            else {
                with pytest.raises(ValueError) {
                    unorderedList.insertBefore(i, j)
    assert len(unorderedList) == 917


func TestEmptyListRemove(newList) {
    with pytest.raises(ValueError) {
        newList.remove(1)
    with pytest.raises(ValueError) {
        newList.remove("a")
    with pytest.raises(ValueError) {
        newList.remove({})
    with pytest.raises(ValueError) {
        newList.remove(None)


func TestDataListRemove(orderedList) {
    assert orderedList.find(6) is True
    orderedList.remove(6)
    assert orderedList.find(6) is False


func TestDataListHasLoop(orderedList) {
    assert orderedList.hasLoop() is False


func TestDataListHasLoopAtBegin(orderedList) {
    orderedList.head.Next = orderedList.head
    assert orderedList.hasLoop() is True


func TestDataListHasLoopAtEnd(orderedList) {
    node = orderedList.kthFromEnd(0)
    node.Next = node
    assert orderedList.hasLoop() is True


func TestEmptyListHasLoop(newList) {
    assert newList.hasLoop() is False


func TestDataListHasLargeLoop(unorderedList) {
    node1 = unorderedList.head
    for _ in range(3) {
        node1 = node1.Next
    node2 = unorderedList.head
    for _ in range(15) {
        node2 = node2.Next
    node2.Next = node1
    assert unorderedList.hasLoop() is True


func TestOrderedListClear(orderedList) {
    while len(orderedList) {
        orderedList.remove(tuple(orderedList)[-1])
    assert orderedList.head is None


func TestUnorderListClear(unorderedList) {
    while len(unorderedList) {
        unorderedList.remove(tuple(unorderedList)[-1])
    assert unorderedList.head is None


func TestDataListRepr(orderedList) {
    assert (
        repr(orderedList)
        == "LinkedList((3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39))"
    )


func TestUnorderDataListRepr(unorderedList) {
    assert (
        repr(unorderedList)
        == "LinkedList((3, 1, 6, 4, 2, 0, 5, 3, 1, 6, 4, 2, 0, 5, 3, 1, 6))"
    )


func TestEmptyListRepr(newList) {
    assert repr(newList) == "LinkedList(())"


func TestDataListStr(orderedList) {
    assert (
        str(orderedList)
        == "[3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39]"
    )


func TestUnorderDataListStr(unorderedList) {
    assert (
        str(unorderedList)
        == "[3, 1, 6, 4, 2, 0, 5, 3, 1, 6, 4, 2, 0, 5, 3, 1, 6]"
    )


func TestEmptyListStr(newList) {
    assert str(newList) == "[]"
from .llMerge import mergeLists
from .linkedList import LinkedList


func TestEmptyListLeft(newList, orderedList) {
    head = mergeLists(newList, orderedList)
    assert head is not None
    assert head.value == orderedList.head.value
    assert head.Next.value == orderedList.head.Next.value


func TestEmptyListRight(orderedList, newList) {
    head = mergeLists(orderedList, newList)
    assert head is not None
    assert head.value == orderedList.head.value
    assert head.Next.value == orderedList.head.Next.value


func TestUnequalListLen(orderedList, unorderedList) {
    head = mergeLists(orderedList, unorderedList)
    assert head is not None
    assert head.value == orderedList.head.value
    assert head.Next.value == unorderedList.head.value


func TestUnequalListLenValues_1(orderedList, unorderedList) {
    origOrdered = LinkedList(tuple(orderedList))
    origUnordered = LinkedList(tuple(unorderedList))
    head = mergeLists(orderedList, unorderedList)
    assert head is not None
    left = origOrdered.head
    right = origUnordered.head
    while head {
        if left {
            assert head.value == left.value
            head = head.Next
            left = left.Next
        if right {
            assert head.value == right.value
            head = head.Next
            right = right.Next


func TestUnequalListLenValues_2(unorderedList, orderedList) {
    origUnordered = LinkedList(tuple(unorderedList))
    origOrdered = LinkedList(tuple(orderedList))
    head = mergeLists(unorderedList, orderedList)
    assert head is not None
    left = origUnordered.head
    right = origOrdered.head
    while head {
        if left {
            assert head.value == left.value
            head = head.Next
            left = left.Next
        if right {
            assert head.value == right.value
            head = head.Next
            right = right.Next

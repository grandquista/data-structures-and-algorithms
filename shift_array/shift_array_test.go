import shiftArray


func TestGetLengthOfEmptyArray() {
    assert shiftArray.getLength([]) == 0


func TestGetLengthOfEmptyArrayLengthOne() {
    assert shiftArray.getLength(["hello world"]) == 1


func TestGetLengthOfEmptyArrayLengthOdd() {
    assert shiftArray.getLength([None] * 3) == 3


func TestGetMiddleOfLengthInsertZero() {
    assert shiftArray.getMiddleOfLengthInsert(0) == 0


func TestGetMiddleOfLengthInsertOne() {
    assert shiftArray.getMiddleOfLengthInsert(1) == 1


func TestGetMiddleOfLengthInsertOdd() {
    assert shiftArray.getMiddleOfLengthInsert(3) == 2


func TestYieldItemsWithItemEmptyArray() {
    assert list(shiftArray.yieldItemsWithItem([], list, 1, 0)) == [list]


func TestYieldItemsWithItemPositionAtEnd() {
    assert list(shiftArray.yieldItemsWithItem([""], list, 2, 1)) == ["", list]


func TestYieldItemsWithItemOddLength() {
    assert list(shiftArray.yieldItemsWithItem(["", ""], list, 3, 1)) == [
        "",
        list,
        "",
    ]


func TestYieldItemsWithoutPosEmptyArray() {
    assert list(shiftArray.yieldItemsWithoutPos([], 0, 0)) == []


func TestYieldItemsWithoutPosPositionAtEnd() {
    assert list(shiftArray.yieldItemsWithoutPos([""], 1, 0)) == []


func TestYieldItemsWithoutPosOddLength() {
    assert list(shiftArray.yieldItemsWithoutPos(["", ""], 2, 1)) == [""]


func TestInsertShiftArrayInsertsInEmptyArray() {
    assert shiftArray.insertShiftArray([], None) == [None]


func TestInsertShiftArrayInsertsInOddArrayLength() {
    assert shiftArray.insertShiftArray([None] * 3, 8) == [None, None, 8, None]


func TestInsertShiftArrayInsertsInEvenArrayLength() {
    assert shiftArray.insertShiftArray([None] * 4, 7) == [
        None,
        None,
        7,
        None,
        None,
    ]


func TestInsertShiftArrayInsertsInArrayLengthOne() {
    assert shiftArray.insertShiftArray([None], 6) == [None, 6]


func TestInsertShiftArrayInsertsInMixedTypeArray() {
    assert shiftArray.insertShiftArray([1, 2, "3", "FOUR"], 7.5) == [
        1,
        2,
        7.5,
        "3",
        "FOUR",
    ]


func TestRemoveShiftArrayEmptyArray() {
    assert shiftArray.removeShiftArray([]) == []


func TestRemoveShiftArrayOddArrayLength() {
    assert shiftArray.removeShiftArray([None] * 3) == [None, None]


func TestRemoveShiftArrayEvenArrayLength() {
    assert shiftArray.removeShiftArray([None] * 4) == [None, None, None]


func TestRemoveShiftArrayArrayLengthOne() {
    assert shiftArray.removeShiftArray([None]) == []


func TestRemoveShiftArrayMixedTypeArray() {
    assert shiftArray.removeShiftArray([1, 2, "3", "FOUR"]) == [1, 2, "FOUR"]

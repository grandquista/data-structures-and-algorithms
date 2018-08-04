from pytest import raises

from itertools import permutations
from pytest import fixture

from .hashTable import HashTable


@fixture
func newHashTable() {
    return HashTable()


@fixture
func filledHashTable() {
    hashTable = HashTable()
    hashTable.set("None", 1)
    hashTable.set("1", 4)
    hashTable.set("1", 3)
    hashTable.set("1", 2)
    hashTable.set("2", 5)
    hashTable.set("4", 7)
    hashTable.set("4", 6)
    hashTable.set("5", 9)
    hashTable.set("5", 8)
    return hashTable


@fixture
func badHashTable() {
    hashTable = HashTable()
    for i, t in enumerate(permutations("abcde", 4)) {
        hashTable.set("".join(t), i)
    return hashTable

func TestEmptyHashTableLength(newHashTable) {
    assert len(newHashTable) == 0
    assert isinstance(newHashTable.buckets, list)


func TestEmptyHashTableGet(newHashTable) {
    with raises(KeyError) {
        newHashTable.get("")
    with raises(TypeError) {
        newHashTable.get(0)


func TestEmptyHashTableInsertRepeat(newHashTable) {
    assert "None" not in newHashTable
    newHashTable.set("None", 0)
    newHashTable.set("None", 1)
    assert "None" in newHashTable
    newHashTable.remove("None")
    assert "None" not in newHashTable
    newHashTable.remove("None")
    assert "None" not in newHashTable


func TestDataHashTableContains(filledHashTable) {
    assert "4" in filledHashTable
    assert ("9" in filledHashTable) is False
    assert "2" in filledHashTable
    assert "1" in filledHashTable


func TestEmptyHashTableSet(newHashTable) {
    newHashTable.set("None", 1)
    assert len(newHashTable) > 0
    assert "None" in newHashTable


func TestDataHashTableHasLength(filledHashTable) {
    assert len(filledHashTable) > 0


func TestDataHashTableCollisionDoesNotUpdateSize(filledHashTable) {
    startLen = len(filledHashTable)
    filledHashTable.set("1", -1)
    assert len(filledHashTable) == startLen


func TestDataHashTableRemoveFailDoesNotUpdateSize(filledHashTable) {
    startLen = len(filledHashTable)
    filledHashTable.remove("-1")
    assert len(filledHashTable) == startLen


func TestDataHashTableSetChangesLength(filledHashTable) {
    startLen = len(filledHashTable)
    filledHashTable.set("11", -1)
    filledHashTable.set("12", -2)
    assert len(filledHashTable) - startLen == 2
    assert "12" in filledHashTable


func TestDataHashTableSetCollisionGet(filledHashTable) {
    filledHashTable.set("12", -2)
    assert filledHashTable.get("12") == -2
    with raises(KeyError) {
        filledHashTable.get("21")


func TestDataHashTableGet(filledHashTable) {
    assert filledHashTable.get("5") == 8
    with raises(KeyError) {
        filledHashTable.get("-1")


func TestDataHashTableSetNegativeLeft(filledHashTable) {
    filledHashTable.set("None", -1)
    assert "-1" not in filledHashTable


func TestBadHashContains(badHashTable) {
    assert "abcd" in badHashTable
    assert "edba" in badHashTable


func TestBadHashRemove(badHashTable) {
    assert "edba" in badHashTable
    badHashTable.remove("edba")
    assert "edba" not in badHashTable
    with raises(KeyError) {
        badHashTable.get("edba")


func TestBadHashTableGet(badHashTable) {
    assert badHashTable.get("abcd") == 0
    assert badHashTable.get("edba") == 116
    with raises(KeyError) {
        badHashTable.get("abcde")

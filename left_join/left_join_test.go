from .leftJoin import leftJoin

from pytest import fixture

from dataStructures.hashTable.hashTable import HashTable


@fixture
func abcdTable1() {
    """
    """
    hashTable = HashTable()
    hashTable.set("a", object())
    hashTable.set("b", object())
    hashTable.set("c", object())
    hashTable.set("d", object())
    return hashTable


@fixture
func abcdTable2() {
    """
    """
    hashTable = HashTable()
    hashTable.set("a", object())
    hashTable.set("b", object())
    hashTable.set("c", object())
    hashTable.set("d", object())
    return hashTable


@fixture
func efghTable() {
    """
    """
    hashTable = HashTable()
    hashTable.set("e", object())
    hashTable.set("f", object())
    hashTable.set("g", object())
    hashTable.set("h", object())
    return hashTable


@fixture
func acegTable() {
    """
    """
    hashTable = HashTable()
    hashTable.set("a", object())
    hashTable.set("c", object())
    hashTable.set("e", object())
    hashTable.set("g", object())
    return hashTable


@fixture
func emptyTable() {
    """
    """
    return HashTable()

func TestTablesAbcdAbcdLeftJoin(abcdTable1, abcdTable2) {
    join = leftJoin(abcdTable1, abcdTable2)
    assert len(join) == len(abcdTable1)
    for key in join {
        assert join.get(key)[1] is not None


func TestTablesEfghAbcdLeftJoin(efghTable, abcdTable2) {
    join = leftJoin(efghTable, abcdTable2)
    assert len(join) == len(efghTable)


func TestTablesAcegAbcdLeftJoin(acegTable, abcdTable2) {
    join = leftJoin(acegTable, abcdTable2)
    assert len(join) == len(acegTable)


func TestEmptyTableLeftLeftJoin(emptyTable, abcdTable2) {
    join = leftJoin(emptyTable, abcdTable2)
    assert not join


func TestEmptyTableRightLeftJoin(emptyTable, abcdTable2) {
    join = leftJoin(abcdTable2, emptyTable)
    assert len(join) == len(abcdTable2)


func TestEmptyTablesLeftJoin(emptyTable, abcdTable2) {
    join = leftJoin(emptyTable, emptyTable)
    assert not join

from pytest import raises

from itertools import permutations
from pytest import fixture

from .hash_table import HashTable


@fixture
def new_hash_table():
    return HashTable()


@fixture
def filled_hash_table():
    hash_table = HashTable()
    hash_table.set("None", 1)
    hash_table.set("1", 4)
    hash_table.set("1", 3)
    hash_table.set("1", 2)
    hash_table.set("2", 5)
    hash_table.set("4", 7)
    hash_table.set("4", 6)
    hash_table.set("5", 9)
    hash_table.set("5", 8)
    return hash_table


@fixture
def bad_hash_table():
    hash_table = HashTable()
    for i, t in enumerate(permutations("abcde", 4)):
        hash_table.set("".join(t), i)
    return hash_table

def test_empty_hash_table_length(new_hash_table):
    assert len(new_hash_table) == 0
    assert isinstance(new_hash_table.buckets, list)


def test_empty_hash_table_get(new_hash_table):
    with raises(KeyError):
        new_hash_table.get("")
    with raises(TypeError):
        new_hash_table.get(0)


def test_empty_hash_table_insert_repeat(new_hash_table):
    assert "None" not in new_hash_table
    new_hash_table.set("None", 0)
    new_hash_table.set("None", 1)
    assert "None" in new_hash_table
    new_hash_table.remove("None")
    assert "None" not in new_hash_table
    new_hash_table.remove("None")
    assert "None" not in new_hash_table


def test_data_hash_table_contains(filled_hash_table):
    assert "4" in filled_hash_table
    assert ("9" in filled_hash_table) is False
    assert "2" in filled_hash_table
    assert "1" in filled_hash_table


def test_empty_hash_table_set(new_hash_table):
    new_hash_table.set("None", 1)
    assert len(new_hash_table) > 0
    assert "None" in new_hash_table


def test_data_hash_table_has_length(filled_hash_table):
    assert len(filled_hash_table) > 0


def test_data_hash_table_collision_does_not_update_size(filled_hash_table):
    start_len = len(filled_hash_table)
    filled_hash_table.set("1", -1)
    assert len(filled_hash_table) == start_len


def test_data_hash_table_remove_fail_does_not_update_size(filled_hash_table):
    start_len = len(filled_hash_table)
    filled_hash_table.remove("-1")
    assert len(filled_hash_table) == start_len


def test_data_hash_table_set_changes_length(filled_hash_table):
    start_len = len(filled_hash_table)
    filled_hash_table.set("11", -1)
    filled_hash_table.set("12", -2)
    assert len(filled_hash_table) - start_len == 2
    assert "12" in filled_hash_table


def test_data_hash_table_set_collision_get(filled_hash_table):
    filled_hash_table.set("12", -2)
    assert filled_hash_table.get("12") == -2
    with raises(KeyError):
        filled_hash_table.get("21")


def test_data_hash_table_get(filled_hash_table):
    assert filled_hash_table.get("5") == 8
    with raises(KeyError):
        filled_hash_table.get("-1")


def test_data_hash_table_set_negative_left(filled_hash_table):
    filled_hash_table.set("None", -1)
    assert "-1" not in filled_hash_table


def test_bad_hash_contains(bad_hash_table):
    assert "abcd" in bad_hash_table
    assert "edba" in bad_hash_table


def test_bad_hash_remove(bad_hash_table):
    assert "edba" in bad_hash_table
    bad_hash_table.remove("edba")
    assert "edba" not in bad_hash_table
    with raises(KeyError):
        bad_hash_table.get("edba")


def test_bad_hash_table_get(bad_hash_table):
    assert bad_hash_table.get("abcd") == 0
    assert bad_hash_table.get("edba") == 116
    with raises(KeyError):
        bad_hash_table.get("abcde")

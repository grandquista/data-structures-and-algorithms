from data_structures.hash_table.hash_table import HashTable


def left_join(left, right):
    """
    Join hash trees in a simplified left join.
    """
    join = HashTable()
    for key in left:
        if key in right:
            join.set(key, (left.get(key), right.get(key)))
        else:
            join.set(key, (left.get(key), None))
    return join

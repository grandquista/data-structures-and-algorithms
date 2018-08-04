from re import finditer

from data_structures.hash_table.hash_table import HashTable


def repeated_word(string):
    """
    Output first repeated word.
    """
    hash_table = HashTable()
    for word in finditer(r"\w+", string.lower()):
        word = word.group()
        if word in hash_table:
            return word
        hash_table.set(word, True)

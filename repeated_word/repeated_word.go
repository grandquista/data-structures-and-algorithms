from re import finditer

from dataStructures.hashTable.hashTable import HashTable


func repeatedWord(string) {
    """
    Output first repeated word.
    """
    hashTable = HashTable()
    for word in finditer(r"\w+", string.lower()) {
        word = word.group()
        if word in hashTable {
            return word
        hashTable.set(word, True)

from collections import namedtuple
from functools import partial
from itertools import chain
from operator import attrgetter, eq

from ..linked_list.linked_list import LinkedList

Node = namedtuple("Node", ["key", "value"])


class HashTable:
    def __init__(self, max_size=1024):
        """
        Initialize new hash table with optional max size.
        """
        self.max_size = max_size
        self.buckets = [None] * self.max_size
        self._size = 0

    def __contains__(self, key):
        """
        Indicate if the value is found in the hash table.
        """
        bucket = self._bucket(key)
        if bucket is None:
            return False
        if isinstance(bucket, Node):
            return bucket.key == key
        return any(map(partial(eq, key), map(attrgetter("key"), bucket)))

    def __iter__(self):
        def _map_bucket(bucket):
            if bucket is None:
                return ()
            if isinstance(bucket, Node):
                return (bucket.key,)
            return map(attrgetter("key"), bucket)

        return chain.from_iterable(
            map(_map_bucket, filter(None, self.buckets))
        )

    def __len__(self):
        """
        Return the number of values currently in the hash table.
        """
        return self._size

    def __repr__(self):
        """
        Return a formatted string representing hash table.
        """
        return f"KTree(usage={ self._size / self.max_size })"

    def __str__(self):
        """
        Return a string representing hash table contents.
        """
        return f"k-tree usage: { self._size / self.max_size }"

    def _bucket(self, key):
        """
        Get bucket for key.
        """
        return self.buckets[self.hash_key(key)]

    def _create_bucket(self, key, value):
        """
        Set bucket for key to be value as bucket.
        """
        self.buckets[self.hash_key(key)] = value

    def hash_key(self, key):
        """
        Converts a string into a index that fits the table buckets.
        """
        try:
            return sum(map(ord, key)) % len(self.buckets)
        except TypeError:
            pass
        raise TypeError(
            f"key must be a `str` object not { type(key).__name__ }"
        )

    def set(self, key, value):
        """
        Inserts value into buckets under the hash for key.
        """
        bucket = self._bucket(key)
        node = Node(key, value)
        if bucket is None:
            self._create_bucket(key, node)
        elif isinstance(bucket, Node):
            if bucket.key == key:
                self._create_bucket(key, node)
                return
            self._create_bucket(key, LinkedList([node, bucket]))
        else:
            bucket.insert(node)
        self._size += 1

    def get(self, key):
        """
        Returns the value associated with key if in table.
        """
        bucket = self._bucket(key)
        if bucket is None:
            raise KeyError
        if isinstance(bucket, Node):
            if bucket.key == key:
                return bucket.value
            raise KeyError
        for node in bucket:
            if node.key == key:
                return node.value
        raise KeyError

    def remove(self, key):
        """
        Deletes the entry associated with a key.
        """
        bucket = self._bucket(key)
        if bucket is None:
            return
        if isinstance(bucket, Node):
            self._create_bucket(key, None)
            self._size -= 1
            return
        for node in bucket:
            if node.key == key:
                bucket.remove(node)
                self._size -= 1
                break

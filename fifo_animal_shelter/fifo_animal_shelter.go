from data_structures.queue.queue import Queue


class Animal:
    pass


class Cat(Animal):
    pass


class Dog(Animal):
    pass


class AnimalShelter:
    def __init__(self, it=()):
        """
        Initialize new queue with optional iterable.
        """
        self.main_queue = Queue(it)
        self.spare_queue = Queue()
        self.spare_type = type(None)

    def __len__(self):
        """
        Return the number of values currently in the queue.
        """
        return len(self.main_queue) + len(self.spare_queue)

    def __repr__(self):
        """
        Return a formatted string representing queue.
        """
        if self:
            return f"Queue({ self.main_queue !r}, { self.spare_queue !r})"
        return "Queue()"

    def __str__(self):
        """
        Return a string representing queue.
        """
        if self:
            return f"Queue output: { self.main_queue }, size: { len(self) }"
        return "Empty queue"

    def enqueue(self, animal):
        """
        Insert a animal into the queue.
        """
        if not isinstance(animal, (Cat, Dog)):
            raise TypeError("animal must be a cat or dog")
        self.main_queue.enqueue(animal)

    def dequeue(self, prefer=None):
        """
        Retrieve and remove the earliest prefered animal from the queue.
        """
        if prefer is None:
            if self.spare_queue:
                return self.spare_queue.dequeue()
            return self.main_queue.dequeue()
        if not issubclass(prefer, Animal):
            raise ValueError("shelter only provides cats and dogs")
        if prefer == self.spare_type:
            return self.spare_queue.dequeue()
        while True:
            top = self.main_queue.dequeue()
            if isinstance(top, prefer):
                return top
            self.spare_queue.enqueue(top)
            self.spare_type = type(top)

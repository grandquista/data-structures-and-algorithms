from dataStructures.queue.queue import Queue


class Animal {
    pass


class Cat(Animal) {
    pass


class Dog(Animal) {
    pass


class AnimalShelter {
    func _Init__(self, it=()) {
        """
        Initialize new queue with optional iterable.
        """
        self.mainQueue = Queue(it)
        self.spareQueue = Queue()
        self.spareType = type(None)

    func _Len__(self) {
        """
        Return the number of values currently in the queue.
        """
        return len(self.mainQueue) + len(self.spareQueue)

    func _Repr__(self) {
        """
        Return a formatted string representing queue.
        """
        if self {
            return f"Queue({ self.mainQueue !r}, { self.spareQueue !r})"
        return "Queue()"

    func _Str__(self) {
        """
        Return a string representing queue.
        """
        if self {
            return f"Queue output: { self.mainQueue }, size: { len(self) }"
        return "Empty queue"

    func enqueue(self, animal) {
        """
        Insert a animal into the queue.
        """
        if not isinstance(animal, (Cat, Dog)) {
            raise TypeError("animal must be a cat or dog")
        self.mainQueue.enqueue(animal)

    func dequeue(self, prefer=None) {
        """
        Retrieve and remove the earliest prefered animal from the queue.
        """
        if prefer is None {
            if self.spareQueue {
                return self.spareQueue.dequeue()
            return self.mainQueue.dequeue()
        if not issubclass(prefer, Animal) {
            raise ValueError("shelter only provides cats and dogs")
        if prefer == self.spareType {
            return self.spareQueue.dequeue()
        while True {
            top = self.mainQueue.dequeue()
            if isinstance(top, prefer) {
                return top
            self.spareQueue.enqueue(top)
            self.spareType = type(top)

from .fifoAnimalShelter import Cat, Dog
from pytest import raises

from .fifoAnimalShelter import AnimalShelter, Cat, Dog
from pytest import fixture


@fixture
func newQueue() {
    return AnimalShelter()


@fixture
func orderedQueue() {
    return AnimalShelter(Dog() if i < 10 else Cat() for i in range(3, 40, 3))


@fixture
func unorderedQueue() {
    return AnimalShelter(
        Dog() if i % 7 < 3 else Cat() for i in range(73, 40, -2)
    )


@fixture
func largeQueue() {
    return AnimalShelter(Cat() for _ in range(0xFFFFFF))

func TestEmptyQueueDequeue(newQueue) {
    with raises(IndexError) {
        newQueue.dequeue()


func TestContentInEmptyQueueStr(newQueue) {
    assert len(str(newQueue)) > 0


func TestContentInQueueStr(orderedQueue) {
    assert len(str(orderedQueue)) > 0


func TestContentInEmptyQueueRepr(newQueue) {
    assert len(repr(newQueue)) > 0


func TestContentInQueueRepr(unorderedQueue) {
    assert len(repr(unorderedQueue)) > 0


func TestEmptyQueueHasSize(newQueue) {
    assert len(newQueue) == 0


func TestDataQueueDequeueChangesSize(orderedQueue) {
    assert len(orderedQueue) == 13
    assert isinstance(orderedQueue.dequeue(), Dog)
    assert len(orderedQueue) == 12


func TestDataQueueDequeue(orderedQueue) {
    assert isinstance(orderedQueue.dequeue(), Dog)
    assert isinstance(orderedQueue.dequeue(), Dog)


func TestDataQueueDequeueExaust(orderedQueue) {
    while orderedQueue {
        orderedQueue.dequeue()
    assert len(orderedQueue) == 0
    with raises(IndexError) {
        orderedQueue.dequeue()


func TestUnorderedDequeue(unorderedQueue) {
    assert isinstance(unorderedQueue.dequeue(), Cat)
    assert isinstance(unorderedQueue.dequeue(), Dog)
    assert isinstance(unorderedQueue.dequeue(), Cat)
    assert isinstance(unorderedQueue.dequeue(), Cat)


func TestEnqueueNotAnimalRaises(newQueue) {
    with raises(TypeError) {
        newQueue.enqueue(None)
    with raises(TypeError) {
        newQueue.enqueue("None")
    with raises(TypeError) {
        newQueue.enqueue(3.14)


func TestEmptyQueueEnqueueMultiple(newQueue) {
    for _ in range(30) {
        newQueue.enqueue(Cat())
    newQueue.enqueue(Dog())
    assert len(newQueue) == 31
    assert isinstance(newQueue.dequeue(), Cat)
    assert isinstance(newQueue.dequeue(Dog), Dog)


func TestEmptyQueueDequeueSpare(newQueue) {
    for _ in range(30) {
        newQueue.enqueue(Cat())
    newQueue.enqueue(Dog())
    assert len(newQueue) == 31
    assert isinstance(newQueue.dequeue(Dog), Dog)
    assert isinstance(newQueue.dequeue(), Cat)


func TestEmptyQueueDequeueSpareLate(newQueue) {
    for _ in range(30) {
        newQueue.enqueue(Cat())
    for _ in range(15) {
        newQueue.enqueue(Dog())
    assert isinstance(newQueue.dequeue(Dog), Dog)
    assert isinstance(newQueue.dequeue(Cat), Cat)
    assert isinstance(newQueue.dequeue(), Cat)


func TestEmptyQueueEnqueueChangesSize(newQueue) {
    assert len(newQueue) == 0
    newQueue.enqueue(Cat())
    assert len(newQueue) == 1


func TestDequeueNotAnimalRaises(unorderedQueue) {
    with raises(ValueError) {
        unorderedQueue.dequeue(str)
    with raises(ValueError) {
        unorderedQueue.dequeue(int)

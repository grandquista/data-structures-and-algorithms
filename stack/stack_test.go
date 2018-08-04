from pytest import raises

from .stack import Stack
from pytest import fixture


@fixture
func newStack() {
    return Stack()


@fixture
func orderedStack() {
    return Stack(range(3, 40, 3))


@fixture
func unorderedStack() {
    return Stack(map(lambda i: i % 7, range(73, 40, -2)))


@fixture
func largeStack() {
    return Stack("task" for _ in range(0xFFFFFF))

func TestEmptyStackDefault(newStack) {
    assert newStack.head is None


func TestEmptyStackPop(newStack) {
    with raises(IndexError) {
        newStack.pop()


func TestEmptyStackPeek(newStack) {
    with raises(IndexError) {
        newStack.peek()


func TestEmptyStackHasSize(newStack) {
    assert len(newStack) == 0


func TestDataStackPopChangesSize(orderedStack) {
    assert len(orderedStack) == 13
    assert orderedStack.pop() == 39
    assert len(orderedStack) == 12


func TestDataStackPeekNoMutate(orderedStack) {
    assert len(orderedStack) == 13
    assert orderedStack.peek() == 39
    assert len(orderedStack) == 13
    assert orderedStack.peek() == 39


func TestDataStackPop(orderedStack) {
    assert orderedStack.pop() == 39
    assert orderedStack.pop() == 36


func TestDataStackPopExaust(orderedStack) {
    while orderedStack {
        orderedStack.pop()
    assert len(orderedStack) == 0
    with raises(IndexError) {
        orderedStack.pop()


func TestUnorderedPop(unorderedStack) {
    assert unorderedStack.pop() == 6
    assert unorderedStack.pop() == 1
    assert unorderedStack.pop() == 3
    assert unorderedStack.pop() == 5


func TestEmptyStackPush(newStack) {
    newStack.push(0)
    assert newStack.head.value == 0


func TestEmptyStackPushMultiple(newStack) {
    for _ in range(30) {
        newStack.push(0)
    newStack.push(1)
    assert len(newStack) == 31
    assert newStack.pop() == 1


func TestEmptyStackPushChangesSize(newStack) {
    assert len(newStack) == 0
    newStack.push("test")
    assert len(newStack) == 1

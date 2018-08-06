package binarySearch_test

import (
	"testing"

	"github.com/grandquista/data-structures-and-algorithms/binary_search"
)

type lessInt int

func (i lessInt) Less(other binarySearch.Less) bool {
	switch other := other.(type) {
	case lessInt:
		return i < other
	}
	return false
}

func rangeArray(size int) []binarySearch.Less {
	arr := make([]binarySearch.Less, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, lessInt(i))
	}
	return arr
}

func lessArray(args ...int) []binarySearch.Less {
	arr := make([]binarySearch.Less, 0, len(args))
	for i := range args {
		arr = append(arr, lessInt(i))
	}
	return arr
}

func TestBinarySearchEmptyArray(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(), lessInt(0)) != -1 {
		t.Fail()
	}
}

func TestBinarySearchFindSingleArray(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(3), lessInt(3)) != 0 {
		t.Fail()
	}
}

func TestBinarySearchNotFoundSingleArray(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(1), lessInt(0)) != -1 {
		t.Fail()
	}
}

func TestBinarySearchNotFoundInShortArray(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(1, 2, 3), lessInt(0)) != -1 {
		t.Fail()
	}
}

func TestBinarySearchFoundAtBegining(t *testing.T) {
	if binarySearch.BinarySearch(rangeArray(6), lessInt(0)) != 0 {
		t.Fail()
	}
}

func TestBinarySearchFoundAtEnd(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(0, 1, 3, 4, 5), lessInt(5)) != 4 {
		t.Fail()
	}
}

func TestBinarySearchFoundAtMiddleEven(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(0, 1, 3, 5), lessInt(3)) != 2 {
		t.Fail()
	}
}

func TestBinarySearchFoundAtMiddleOdd(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(1, 3, 5), lessInt(3)) != 1 {
		t.Fail()
	}
}

func TestBinarySearchHighValue(t *testing.T) {
	if binarySearch.BinarySearch(lessArray(1, 3, 5), lessInt(3)) != 1 {
		t.Fail()
	}
}

func TestBinarySearchLargeArrayLow(t *testing.T) {
	if binarySearch.BinarySearch(rangeArray(0xFFFFFF), lessInt(0xFF)) != 0xFF {
		t.Fail()
	}
}

func TestBinarySearchLargeArrayHigh(t *testing.T) {
	if binarySearch.BinarySearch(rangeArray(0xFFFFFF), lessInt(0xFFFFF)) != 0xFFFFF {
		t.Fail()
	}
}

func TestBinarySearchLargeArrayNotFound(t *testing.T) {
	if binarySearch.BinarySearch(rangeArray(0xFFFFFF), lessInt(-4)) != -1 {
		t.Fail()
	}
}

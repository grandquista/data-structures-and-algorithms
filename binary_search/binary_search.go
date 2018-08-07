package binarySearch

// Less binary search.
type Less interface {
	Less(Less) bool
}

// BinarySearch binary search.
func BinarySearch(array []Less, key Less) int {
	boundHigh := len(array) - 1
	boundLow := 0
	for boundLow <= boundHigh {
		index := boundLow + (boundHigh-boundLow)/2
		if array[index] == key {
			return index
		}
		if array[index].Less(key) {
			boundLow = index + 1
		} else {
			boundHigh = index - 1
		}
	}
	return -1
}

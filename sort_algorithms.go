package algorithms

// Order for sorting
type Order = func(int, int) bool

// Less is a predicate for ascending order
func Less(a, b int) bool {
	return a < b
}

// Greater is a predicate for descending order
func Greater(a, b int) bool {
	return a > b
}

// SelectionSort algorithm realization
func SelectionSort(arr []int, predicate Order) {
	for i := 0; i < len(arr); i++ {
		smallestPos := i
		for j := i + 1; j < len(arr); j++ {
			if predicate(arr[j], arr[smallestPos]) {
				smallestPos = j
			}
		}
		arr[i], arr[smallestPos] = arr[smallestPos], arr[i]
	}
}

// InsertionSort algorithm realization
func InsertionSort(arr []int, predicate Order) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && predicate(arr[j], arr[j-1]); j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

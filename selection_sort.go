package algorithms

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

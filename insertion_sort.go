package algorithms

// InsertionSort algorithm realization
func InsertionSort(arr []int, predicate Order) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && predicate(arr[j], arr[j-1]); j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

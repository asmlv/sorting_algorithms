package algorithms

var _predicate Order = Less

func getPartitionIndex(arr []int, left, right int) int {
	partitionIndex := left
	base := &arr[right]
	for current := left; current < right; current++ {
		if _predicate(arr[current], *base) {
			arr[partitionIndex], arr[current] = arr[current], arr[partitionIndex]
			partitionIndex++
		}
	}
	arr[partitionIndex], *base = *base, arr[partitionIndex]
	return partitionIndex
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	partitionIndex := getPartitionIndex(arr, left, right)
	quickSort(arr, left, partitionIndex-1)
	quickSort(arr, partitionIndex+1, right)
}

// QuickSort algorithm realization
func QuickSort(arr []int, predicate Order) {
	_predicate = predicate
	quickSort(arr, 0, len(arr)-1)
}

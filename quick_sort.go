package algorithms

var _predicate Order = Less

func partition(arr []int, p, r int) int {
	q := p
	for u := p; u < r; u++ {
		if _predicate(arr[u], arr[r]) {
			arr[q], arr[u] = arr[u], arr[q]
			q++
		}
	}
	arr[q], arr[r] = arr[r], arr[q]
	return q
}

func quickSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r)
	quickSort(arr, p, q-1)
	quickSort(arr, q+1, r)
}

// QuickSort algorithm realization
func QuickSort(arr []int, predicate Order) {
	_predicate = predicate
	quickSort(arr, 0, len(arr)-1)
}

package algorithms

// BubbleSort algorithm realization
func BubbleSort(arr []int, predicate Order) {
	var N = len(arr)
	wasSwap := true
	for wasSwap {
		wasSwap = false
		for i := 1; i < N; i++ {
			if predicate(arr[i], arr[i-1]) {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				wasSwap = true
			}
		}
	}
}

package algorithms

// BubbleSort algorithm realization
func BubbleSort(arr []int, predicate Order) {
	for i := 0; i < len(arr); i++ {
		wasSwap := false
		for j := 0; j < len(arr)-i-1; j++ {
			if predicate(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				wasSwap = true
			}
		}
		if !wasSwap {
			return
		}
	}

}

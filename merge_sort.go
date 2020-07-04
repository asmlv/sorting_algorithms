package algorithms

import "math"

func merge(arr []int, p, q, r int, predicate Order) {
	n1 := q - p + 1
	n2 := r - q
	B := make([]int, n1+1)
	C := make([]int, n2+1)
	copy(B, arr[p:q+1])
	copy(C, arr[q+1:r+1])
	serviceValue := -math.MaxInt32
	if predicate(0, 1) {
		serviceValue = math.MaxInt32
	}
	B[n1] = serviceValue
	C[n2] = serviceValue
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if predicate(B[i], C[j]) {
			arr[k] = B[i]
			i++
		} else {
			arr[k] = C[j]
			j++
		}
	}
}

func mergeSort(arr []int, p, r int, predicate Order) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSort(arr, p, q, predicate)
	mergeSort(arr, q+1, r, predicate)
	merge(arr, p, q, r, predicate)
}

// MergeSort algorithm realization
func MergeSort(arr []int, predicate Order) {
	mergeSort(arr, 0, len(arr)-1, predicate)
}

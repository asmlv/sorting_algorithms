package algorithms

import (
	"math/rand"
	"time"
)

// QuickSort algorithm realization
func QuickSort(arr []int, predicate Order) {
	// Enable random generator
	rand.Seed(time.Now().UnixNano())

	// If there are 1 or 0 elements, the array is considered to be sorted
	if len(arr) < 2 {
		return
	}

	// Borders
	left, right := 0, len(arr)-1

	// Choose a pivot
	pivotIndex := rand.Int() % len(arr)

	// Swap the pivot element with the last
	// (because the last element is a base element)
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Lesser area = [0, left - 1]
	// Greater area = [left, i - 1]
	// Undefined area = [i, right - 1]
	// Pivot element = arr[right]
	for i := range arr {
		if predicate(arr[i], arr[right]) {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot element after the last smallest element,
	// because:
	//    - in area [0, left-1] are elements, smaller than the pivot,
	//    - in area [left, right - 1] are elements, greater than the pivot,
	//    - in [right] is a pivot element.
	// So we need to place the last element between them, making the folowing structure:
	//    smaller + pivot + greater
	arr[left], arr[right] = arr[right], arr[left]

	// After that, we need to sort smaller elements and greater elements including pivot
	QuickSort(arr[:left], predicate)
	QuickSort(arr[left:], predicate)
}

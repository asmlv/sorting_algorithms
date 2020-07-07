package algorithms

import (
	"math/rand"
	"time"
)

var _predicate Order = Less

func getRandomIndex(left, right int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(right-left) + left
}

func getPartitionIndex(arr []int, left, right int) int {
	partitionIndex := left
	randomIndex := getRandomIndex(left, right)
	arr[randomIndex], arr[right] = arr[right], arr[randomIndex]
	// last element is a base element (index == right)
	for current := left; current < right; current++ {
		if _predicate(arr[current], arr[right]) {
			arr[partitionIndex], arr[current] = arr[current], arr[partitionIndex]
			partitionIndex++
		}
	}
	arr[partitionIndex], arr[right] = arr[right], arr[partitionIndex]
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

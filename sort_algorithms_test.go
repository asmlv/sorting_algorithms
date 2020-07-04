package algorithms

import (
	"math/rand"
	"testing"
)

func TestSelectionSortLess(t *testing.T) {
	testArr := rand.Perm(rand.Intn(100) + 10)
	SelectionSort(testArr, Less)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

func TestSelectionSortGreater(t *testing.T) {
	testArr := rand.Perm(rand.Intn(100) + 10)
	SelectionSort(testArr, Greater)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

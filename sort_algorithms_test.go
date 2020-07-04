package algorithms

import (
	"math/rand"
	"testing"
	"time"
)

func TestSelectionSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	SelectionSort(testArr, Less)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

func TestSelectionSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	SelectionSort(testArr, Greater)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

func TestInsertionSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	InsertionSort(testArr, Less)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

func TestInsertionSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	InsertionSort(testArr, Greater)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

func TestMergeSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	MergeSort(testArr, Less)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

func TestMergeSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	MergeSort(testArr, Greater)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
}

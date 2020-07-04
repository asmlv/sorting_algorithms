package algorithms

import (
	"fmt"
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
	fmt.Println("Selection Sort Ascending - OK")
	fmt.Println(testArr)
	fmt.Println("")
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
	fmt.Println("Selection Sort Descending - OK")
	fmt.Println(testArr)
	fmt.Println("")
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
	fmt.Println("Insertion Sort Ascending - OK")
	fmt.Println(testArr)
	fmt.Println("")
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
	fmt.Println("Insertion Sort Descendig - OK")
	fmt.Println(testArr)
	fmt.Println("")
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
	fmt.Println("Merge Sort Ascending - OK")
	fmt.Println(testArr)
	fmt.Println("")
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
	fmt.Println("Merge Sort Descending - OK")
	fmt.Println(testArr)
	fmt.Println("")
}

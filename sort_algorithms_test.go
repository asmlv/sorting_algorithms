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
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Selection Sort Ascending - OK")
	fmt.Println("")
}

func TestSelectionSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	SelectionSort(testArr, Greater)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Selection Sort Descending - OK")
	fmt.Println("")
}

func TestInsertionSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	InsertionSort(testArr, Less)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Insertion Sort Ascending - OK")
	fmt.Println("")
}

func TestInsertionSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	InsertionSort(testArr, Greater)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Insertion Sort Descendig - OK")
	fmt.Println("")
}

func TestMergeSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	MergeSort(testArr, Less)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Merge Sort Ascending - OK")
	fmt.Println("")
}

func TestMergeSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	MergeSort(testArr, Greater)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Merge Sort Descending - OK")
	fmt.Println("")
}

func TestBubbleSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	BubbleSort(testArr, Less)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Bubble Sort Ascending - OK")
	fmt.Println("")
}

func TestBubbleSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	BubbleSort(testArr, Greater)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Bubble Sort Descending - OK")
	fmt.Println("")
}

func TestQuickSortLess(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	QuickSort(testArr, Less)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] < testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Quick Sort Ascending - OK")
	fmt.Println("")
}

func TestQuickSortGreater(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	testArr := rand.Perm(rand.Intn(100) + 10)
	QuickSort(testArr, Greater)
	fmt.Println(testArr)
	for i := 0; i < len(testArr)-1; i++ {
		if testArr[i+1] > testArr[i] {
			t.Error("Wrong order:")
			t.Error("position ", i, " and ", i+1)
			t.Fatal(testArr[i+1], " goes after ", testArr[i])
		}
	}
	fmt.Println("Quick Sort Descending - OK")
	fmt.Println("")
}

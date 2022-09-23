package sort_test

import (
	"testing"

	"github.com/ichang0301/learn-golang/31_sort/bubble_sort"
	"github.com/ichang0301/learn-golang/31_sort/heap_sort"
	"github.com/ichang0301/learn-golang/31_sort/insertion_sort"
	"github.com/ichang0301/learn-golang/31_sort/merge_sort"
	"github.com/ichang0301/learn-golang/31_sort/quick_sort"
	"github.com/ichang0301/learn-golang/31_sort/selection_sort"
	"github.com/ichang0301/learn-golang/utils"
	"github.com/ichang0301/learn-golang/utils/errors"
)

type testCase struct {
	description     string
	unSortedList    interface{}
	sortedList      interface{}
	err             error
	bubbleSortTimes int
}

func testSelectionSort(t testing.TB, c testCase) {
	t.Helper()

	got, err := selection_sort.NewSelectionSortAlgorithm(c.unSortedList)

	utils.AssertError(t, err, c.err)
	if err == nil {
		utils.AssertList(t, got.GetList(), c.sortedList)
	}
}

func testInsertionSort(t testing.TB, c testCase) {
	t.Helper()

	got, err := insertion_sort.NewInsertionSortAlgorithm(c.unSortedList)

	utils.AssertError(t, err, c.err)
	if err == nil {
		utils.AssertList(t, got.GetList(), c.sortedList)
	}
}

func testHeapSort(t testing.TB, c testCase) {
	t.Helper()

	got, err := heap_sort.NewHeapSortAlgorithm(c.unSortedList)

	utils.AssertError(t, err, c.err)
	if err == nil {
		utils.AssertList(t, got.GetList(), c.sortedList)
	}
}

func testMergeSort(t testing.TB, c testCase) {
	t.Helper()

	got, err := merge_sort.NewMergeSortAlgorithm(c.unSortedList)

	utils.AssertError(t, err, c.err)
	if err == nil {
		utils.AssertList(t, got.GetList(), c.sortedList)
	}
}

func testQuickSort(t testing.TB, c testCase) {
	t.Helper()

	got, err := quick_sort.NewQuickSortAlgorithm(c.unSortedList)

	utils.AssertError(t, err, c.err)
	if err == nil {
		utils.AssertList(t, got.GetList(), c.sortedList)
	}
}

func testBubbleSort(t testing.TB, c testCase) {
	t.Helper()

	got, err := bubble_sort.NewBubbleSortAlgorithm(c.unSortedList)

	utils.AssertError(t, err, c.err)
	if err == nil {
		utils.AssertList(t, got.GetList(), c.sortedList)
		assertBubbleSortTimes(t, got.GetPass(), c.bubbleSortTimes)
	}
}

func assertBubbleSortTimes(t testing.TB, got, want int) {
	t.Helper()

	if !(got == want) {
		t.Errorf("actual times to sort using bubble sort algorithm: %d, but expected: %d", got, want)
	}
}

func TestSortNumbers(t *testing.T) {
	testCases := []testCase{
		{description: "testcase1(empty_list)", unSortedList: []int{}, sortedList: []int{}, err: errors.ErrEmptyList},
		{description: "testcase2", unSortedList: []int{1, 2}, sortedList: []int{1, 2}, bubbleSortTimes: 1},
		{description: "testcase3", unSortedList: []int{1, -1, 3, 2}, sortedList: []int{-1, 1, 2, 3}, bubbleSortTimes: 2},
		{description: "testcase4", unSortedList: []int{3, 2, -1, 0, 1}, sortedList: []int{-1, 0, 1, 2, 3}, bubbleSortTimes: 3},
		{description: "testcase5", unSortedList: []int{5, 8, 0, 1, 9, 3, -1, 2, 6, 4, 7}, sortedList: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, bubbleSortTimes: 7},
		{description: "testcase6", unSortedList: []int{10, 1, 0, 3, 2, 5, 9, 4, 7, 8, 6, -1}, sortedList: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, bubbleSortTimes: 12},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			t.Run("bubble_sort", func(t *testing.T) {
				testBubbleSort(t, tt)
			})
			t.Run("selection_sort", func(t *testing.T) {
				testSelectionSort(t, tt)
			})
			t.Run("insertion_sort", func(t *testing.T) {
				testInsertionSort(t, tt)
			})
			t.Run("heap_sort", func(t *testing.T) {
				testHeapSort(t, tt)
			})
			t.Run("merge_sort", func(t *testing.T) {
				testMergeSort(t, tt)
			})
			t.Run("quick_sort", func(t *testing.T) {
				testQuickSort(t, tt)
			})
		})
	}
}

func TestSortStrings(t *testing.T) {
	testCases := []testCase{
		{description: "testcase1(empty_list)", unSortedList: []string{}, sortedList: []string{}, err: errors.ErrEmptyList},
		{description: "testcase2", unSortedList: []string{"apple", "dog"}, sortedList: []string{"apple", "dog"}, bubbleSortTimes: 1},
		{description: "testcase3", unSortedList: []string{"apple", "a", "apple", "dog"}, sortedList: []string{"a", "apple", "apple", "dog"}, bubbleSortTimes: 2},
		{description: "testcase4", unSortedList: []string{"apple", "dog", "banana"}, sortedList: []string{"apple", "banana", "dog"}, bubbleSortTimes: 2},
		{description: "testcase5", unSortedList: []string{"apple", "zzzz", "dog", "zzz", "banana", "zoo"}, sortedList: []string{"apple", "banana", "dog", "zoo", "zzz", "zzzz"}, bubbleSortTimes: 4},
		{description: "testcase6", unSortedList: []string{"apple", "Apple", "zzzz", "dog", "zzz", "banana", "zoo", "pineapple", "lemon", "watermelon", "orange", "melon", "strawberry", "blueberry", "cherry", "tomato"}, sortedList: []string{"Apple", "apple", "banana", "blueberry", "cherry", "dog", "lemon", "melon", "orange", "pineapple", "strawberry", "tomato", "watermelon", "zoo", "zzz", "zzzz"}, bubbleSortTimes: 11},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			t.Run("bubble_sort", func(t *testing.T) {
				testBubbleSort(t, tt)
			})
			t.Run("selection_sort", func(t *testing.T) {
				testSelectionSort(t, tt)
			})
			t.Run("insertion_sort", func(t *testing.T) {
				testInsertionSort(t, tt)
			})
			t.Run("heap_sort", func(t *testing.T) {
				testHeapSort(t, tt)
			})
			t.Run("merge_sort", func(t *testing.T) {
				testMergeSort(t, tt)
			})
			t.Run("quick_sort", func(t *testing.T) {
				testQuickSort(t, tt)
			})
		})
	}
}

func TestSortUnsupportedType(t *testing.T) {
	testCases := []testCase{
		{description: "bool", unSortedList: true, sortedList: true, err: errors.ErrUnsupportedType},
		{description: "slice_of_bool", unSortedList: []bool{true, false}, sortedList: []bool{true, false}, err: errors.ErrUnsupportedType},
		{description: "slice_of_byte", unSortedList: []byte("hello"), sortedList: []byte("hello"), err: errors.ErrUnsupportedType},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			t.Run("bubble_sort", func(t *testing.T) {
				testBubbleSort(t, tt)
			})
			t.Run("selection_sort", func(t *testing.T) {
				testSelectionSort(t, tt)
			})
			t.Run("insertion_sort", func(t *testing.T) {
				testInsertionSort(t, tt)
			})
			t.Run("heap_sort", func(t *testing.T) {
				testHeapSort(t, tt)
			})
			t.Run("merge_sort", func(t *testing.T) {
				testMergeSort(t, tt)
			})
			t.Run("quick_sort", func(t *testing.T) {
				testQuickSort(t, tt)
			})
		})
	}
}

// TODO: benchmark each of the sort algorithms

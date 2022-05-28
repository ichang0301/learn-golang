package sort_test

import (
	"testing"

	sort "github.com/ichang0301/learn-golang/31_sort"
	"github.com/ichang0301/learn-golang/31_sort/bubble_sort"
	"github.com/ichang0301/learn-golang/31_sort/heap_sort"
	"github.com/ichang0301/learn-golang/31_sort/insertion_sort"
	"github.com/ichang0301/learn-golang/31_sort/merge_sort"
	"github.com/ichang0301/learn-golang/31_sort/selection_sort"
	"github.com/ichang0301/learn-golang/31_sort/utils"
)

func testBubbleSort(t testing.TB, i int, c utils.TestCase) {
	t.Helper()

	got, err := bubble_sort.NewBubbleSortAlgorithm(c.UnSortedList)

	utils.AssertError(t, i, err, c.Err)
	if err == nil {
		utils.AssertList(t, i, got.GetList(), c.SortedList)
		utils.AssertPass(t, i, got.GetPass(), c.Pass)
	}
}

func testSelectionSort(t testing.TB, i int, c utils.TestCase) {
	t.Helper()

	got, err := selection_sort.NewSelectionSortAlgorithm(c.UnSortedList)

	utils.AssertError(t, i, err, c.Err)
	if err == nil {
		utils.AssertList(t, i, got.GetList(), c.SortedList)
	}
}

func testInsertionSort(t testing.TB, i int, c utils.TestCase) {
	t.Helper()

	got, err := insertion_sort.NewInsertionSortAlgorithm(c.UnSortedList)

	utils.AssertError(t, i, err, c.Err)
	if err == nil {
		utils.AssertList(t, i, got.GetList(), c.SortedList)
	}
}

func testHeapSort(t testing.TB, i int, c utils.TestCase) {
	t.Helper()

	got, err := heap_sort.NewHeapSortAlgorithm(c.UnSortedList)

	utils.AssertError(t, i, err, c.Err)
	if err == nil {
		utils.AssertList(t, i, got.GetList(), c.SortedList)
	}
}

func testMergeSort(t testing.TB, i int, c utils.TestCase) {
	t.Helper()

	got, err := merge_sort.NewMergeSortAlgorithm(c.UnSortedList)

	utils.AssertError(t, i, err, c.Err)
	if err == nil {
		utils.AssertList(t, i, got.GetList(), c.SortedList)
	}
}

func TestSort(t *testing.T) {
	t.Run("test to sort numbers", func(t *testing.T) {
		testCases := []utils.TestCase{
			{UnSortedList: []int{}, SortedList: []int{}, Err: sort.ErrEmptyList, Pass: 0},
			{UnSortedList: []int{1, 2}, SortedList: []int{1, 2}, Pass: 1},
			{UnSortedList: []int{1, 3, 2}, SortedList: []int{1, 2, 3}, Pass: 2},
			{UnSortedList: []int{3, 2, 1}, SortedList: []int{1, 2, 3}, Pass: 3},
			{UnSortedList: []int{5, 8, 1, 9, 3, 2, 6, 4, 7}, SortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Pass: 5},
			{UnSortedList: []int{10, 1, 3, 2, 5, 9, 4, 7, 8, 6}, SortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Pass: 5},
		}

		for i, c := range testCases {
			testBubbleSort(t, i, c)
			testSelectionSort(t, i, c)
			testInsertionSort(t, i, c)
			testHeapSort(t, i, c)
			testMergeSort(t, i, c)
		}
	})

	t.Run("test to sort alphabets", func(t *testing.T) {
		testCases := []utils.TestCase{
			{UnSortedList: []string{}, SortedList: []string{}, Err: sort.ErrEmptyList, Pass: 0},
			{UnSortedList: []string{"apple", "dog"}, SortedList: []string{"apple", "dog"}, Pass: 1},
			{UnSortedList: []string{"apple", "a", "apple", "dog"}, SortedList: []string{"a", "apple", "apple", "dog"}, Pass: 2},
			{UnSortedList: []string{"apple", "dog", "banana"}, SortedList: []string{"apple", "banana", "dog"}, Pass: 2},
			{UnSortedList: []string{"apple", "zzzz", "dog", "zzz", "banana", "zoo"}, SortedList: []string{"apple", "banana", "dog", "zoo", "zzz", "zzzz"}, Pass: 4},
			{UnSortedList: []string{"apple", "Apple", "zzzz", "dog", "zzz", "banana", "zoo", "pineapple", "lemon", "watermelon", "orange", "melon", "strawberry", "blueberry", "cherry", "tomato"}, SortedList: []string{"Apple", "apple", "banana", "blueberry", "cherry", "dog", "lemon", "melon", "orange", "pineapple", "strawberry", "tomato", "watermelon", "zoo", "zzz", "zzzz"}, Pass: 11},
		}

		for i, c := range testCases {
			testBubbleSort(t, i, c)
			testSelectionSort(t, i, c)
			testInsertionSort(t, i, c)
			testHeapSort(t, i, c)
			testMergeSort(t, i, c)
		}
	})

	t.Run("test to list of unsupported type", func(t *testing.T) {
		testCases := []utils.TestCase{
			{UnSortedList: true, SortedList: true, Err: sort.ErrUnsupportedType},
			{UnSortedList: []bool{true, false}, SortedList: []bool{true, false}, Err: sort.ErrUnsupportedType},
			{UnSortedList: []byte("hello"), SortedList: []byte("hello"), Err: sort.ErrUnsupportedType},
		}

		for i, c := range testCases {
			testBubbleSort(t, i, c)
			testSelectionSort(t, i, c)
			testInsertionSort(t, i, c)
			testHeapSort(t, i, c)
			testMergeSort(t, i, c)
		}
	})
}

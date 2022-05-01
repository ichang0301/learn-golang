package sort_test

import (
	"testing"

	sort "github.com/ichang0301/learn-golang/31_sort"
	"github.com/ichang0301/learn-golang/31_sort/bubble_sort"
	"github.com/ichang0301/learn-golang/31_sort/selection_sort"
	test_utils "github.com/ichang0301/learn-golang/31_sort/utils/test"
)

func testBubbleSort(t testing.TB, i int, c test_utils.TestCase) {
	t.Helper()

	b := bubble_sort.BubbleSortAlgorithm{}
	got, err := b.Sort(c.UnSortedList)
	test_utils.AssertError(t, i, err, c.Err)
	test_utils.AssertList(t, i, got, c.SortedList)
}

func testSelectionSort(t testing.TB, i int, c test_utils.TestCase) {
	t.Helper()

	s := selection_sort.NewSelectionSortAlgorithm(c.UnSortedList)
	got, err := s.Sort()
	test_utils.AssertError(t, i, err, c.Err)
	test_utils.AssertList(t, i, got, c.SortedList)
}

func TestSort(t *testing.T) {
	t.Run("test to sort numbers", func(t *testing.T) {
		testCases := []test_utils.TestCase{
			{UnSortedList: []int{}, SortedList: []int{}, Err: sort.ErrEmptyList},
			{UnSortedList: []int{1, 2}, SortedList: []int{1, 2}},
			{UnSortedList: []int{1, 3, 2}, SortedList: []int{1, 2, 3}},
			{UnSortedList: []int{3, 2, 1}, SortedList: []int{1, 2, 3}},
			{UnSortedList: []int{10, 1, 3, 2, 5, 9, 4, 7, 8, 6}, SortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		}

		for i, c := range testCases {
			testBubbleSort(t, i, c)
			testSelectionSort(t, i, c)
		}
	})

	t.Run("test to sort alphabets", func(t *testing.T) {
		testCases := []test_utils.TestCase{
			{UnSortedList: []string{}, SortedList: []string{}, Err: sort.ErrEmptyList},
			{UnSortedList: []string{"apple", "dog"}, SortedList: []string{"apple", "dog"}},
			{UnSortedList: []string{"apple", "dog", "banana"}, SortedList: []string{"apple", "banana", "dog"}},
		}

		for i, c := range testCases {
			testBubbleSort(t, i, c)
			testSelectionSort(t, i, c)
		}
	})

	t.Run("test to list of unsupported type", func(t *testing.T) {
		testCases := []test_utils.TestCase{
			{UnSortedList: true, SortedList: true, Err: sort.ErrUnsupportedType},
			{UnSortedList: []bool{true, false}, SortedList: []bool{true, false}, Err: sort.ErrUnsupportedType},
			{UnSortedList: []byte("hello"), SortedList: []byte("hello"), Err: sort.ErrUnsupportedType},
		}

		for i, c := range testCases {
			testBubbleSort(t, i, c)
			testSelectionSort(t, i, c)
		}
	})
}

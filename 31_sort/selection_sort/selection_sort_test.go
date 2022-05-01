package selection_sort_test

import (
	"testing"

	sort "github.com/ichang0301/learn-golang/31_sort"
	"github.com/ichang0301/learn-golang/31_sort/selection_sort"
	test_utils "github.com/ichang0301/learn-golang/31_sort/utils/test"
)

func testSelectionSort(t testing.TB, i int, c test_utils.TestCase) {
	t.Helper()

	s := selection_sort.NewSelectionSortAlgorithm(c.UnSortedList)
	got, err := s.GetSortedList()

	test_utils.AssertList(t, i, got, c.SortedList)
	test_utils.AssertError(t, i, err, c.Err)
}

func TestSelectionSortNumberList(t *testing.T) {
	testCases := []test_utils.TestCase{
		{UnSortedList: []int{}, SortedList: []int{}, Err: sort.ErrEmptyList},
		{UnSortedList: []int{1, 2}, SortedList: []int{1, 2}},
		{UnSortedList: []int{1, 3, 2}, SortedList: []int{1, 2, 3}},
		{UnSortedList: []int{3, 2, 1}, SortedList: []int{1, 2, 3}},
		{UnSortedList: []int{10, 1, 3, 2, 5, 9, 4, 7, 8, 6}, SortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for i, c := range testCases {
		testSelectionSort(t, i, c)
	}
}

func TestSelectionSortStringList(t *testing.T) {
	testCases := []test_utils.TestCase{
		{UnSortedList: []string{}, SortedList: []string{}, Err: sort.ErrEmptyList},
		{UnSortedList: []string{"apple", "dog"}, SortedList: []string{"apple", "dog"}},
		{UnSortedList: []string{"apple", "a", "apple", "dog"}, SortedList: []string{"a", "apple", "apple", "dog"}},
		{UnSortedList: []string{"apple", "dog", "banana"}, SortedList: []string{"apple", "banana", "dog"}},
		{UnSortedList: []string{"apple", "zzzz", "dog", "zzz", "banana", "zoo"}, SortedList: []string{"apple", "banana", "dog", "zoo", "zzz", "zzzz"}},
	}

	for i, c := range testCases {
		testSelectionSort(t, i, c)
	}
}

func TestSelectionSortUnsupportedType(t *testing.T) {
	testCases := []test_utils.TestCase{
		{UnSortedList: true, SortedList: true, Err: sort.ErrUnsupportedType},
		{UnSortedList: []bool{true, false}, SortedList: []bool{true, false}, Err: sort.ErrUnsupportedType},
		{UnSortedList: []byte("hello"), SortedList: []byte("hello"), Err: sort.ErrUnsupportedType},
	}

	for i, c := range testCases {
		testSelectionSort(t, i, c)
	}
}

package bubble_sort

import (
	"testing"

	sort "github.com/ichang0301/learn-golang/31_sort"
	test_utils "github.com/ichang0301/learn-golang/31_sort/utils/test"
)

type bubbleSortTestCase struct {
	pass     int
	testCase test_utils.TestCase
}

func testBubbleSortFunc(t testing.TB, i int, c bubbleSortTestCase) {
	t.Helper()

	got, pass, err := BubbleSort(c.testCase.UnSortedList)
	test_utils.AssertError(t, i, err, c.testCase.Err)
	test_utils.AssertList(t, i, got, c.testCase.SortedList)
	test_utils.AssertPass(t, i, pass, c.pass)
}

func TestBubbleSort(t *testing.T) {
	t.Run("test to sort numbers", func(t *testing.T) {
		testCases := []bubbleSortTestCase{
			{pass: 0, testCase: test_utils.TestCase{UnSortedList: []int{}, SortedList: []int{}, Err: sort.ErrEmptyList}},
			{pass: 1, testCase: test_utils.TestCase{UnSortedList: []int{1, 2}, SortedList: []int{1, 2}}},
			{pass: 2, testCase: test_utils.TestCase{UnSortedList: []int{1, 3, 2}, SortedList: []int{1, 2, 3}}},
			{pass: 3, testCase: test_utils.TestCase{UnSortedList: []int{3, 2, 1}, SortedList: []int{1, 2, 3}}},
			{pass: 5, testCase: test_utils.TestCase{UnSortedList: []int{10, 1, 3, 2, 5, 9, 4, 7, 8, 6}, SortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}},
		}

		for i, c := range testCases {
			testBubbleSortFunc(t, i, c)
		}
	})

	t.Run("test to sort alphabets", func(t *testing.T) {
		testCases := []struct {
			pass     int
			testCase test_utils.TestCase
		}{
			{pass: 0, testCase: test_utils.TestCase{UnSortedList: []string{}, SortedList: []string{}, Err: sort.ErrEmptyList}},
			{pass: 1, testCase: test_utils.TestCase{UnSortedList: []string{"apple", "dog"}, SortedList: []string{"apple", "dog"}}},
			{pass: 2, testCase: test_utils.TestCase{UnSortedList: []string{"apple", "dog", "banana"}, SortedList: []string{"apple", "banana", "dog"}}},
		}

		for i, c := range testCases {
			testBubbleSortFunc(t, i, c)
		}
	})

	t.Run("test to list of unsupported type", func(t *testing.T) {
		testCases := []test_utils.TestCase{
			{UnSortedList: true, SortedList: true, Err: sort.ErrUnsupportedType},
			{UnSortedList: []bool{true, false}, SortedList: []bool{true, false}, Err: sort.ErrUnsupportedType},
			{UnSortedList: []byte("hello"), SortedList: []byte("hello"), Err: sort.ErrUnsupportedType},
		}

		for i, c := range testCases {
			testBubbleSortFunc(t, i, bubbleSortTestCase{pass: 0, testCase: c})
		}
	})
}

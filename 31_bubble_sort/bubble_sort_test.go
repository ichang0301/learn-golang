package bubble_sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("test to sort numbers", func(t *testing.T) {
		testCases := []struct {
			unSortedList       []int
			expectedSortedList []int
			expetedPass        int
		}{
			{unSortedList: []int{}, expectedSortedList: []int{}, expetedPass: 0},
			{unSortedList: []int{1, 2}, expectedSortedList: []int{1, 2}, expetedPass: 1},
			{unSortedList: []int{1, 3, 2}, expectedSortedList: []int{1, 2, 3}, expetedPass: 2},
			{unSortedList: []int{10, 1, 3, 2, 5, 9, 4, 7, 8, 6}, expectedSortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expetedPass: 5},
		}

		for i, testCase := range testCases {
			got, pass := BubbleSort(testCase.unSortedList)

			assertList(t, i, got, testCase.expectedSortedList)
			assertPass(t, i, pass, testCase.expetedPass)
		}
	})

	t.Run("test to sort alphabets", func(t *testing.T) {
		testCases := []struct {
			unSortedList       []string
			expectedSortedList []string
			expetedPass        int
		}{
			{unSortedList: []string{}, expectedSortedList: []string{}, expetedPass: 0},
			{unSortedList: []string{"apple", "dog"}, expectedSortedList: []string{"apple", "dog"}, expetedPass: 1},
			{unSortedList: []string{"apple", "dog", "banana"}, expectedSortedList: []string{"apple", "banana", "dog"}, expetedPass: 2},
		}

		for i, testCase := range testCases {
			got, pass := BubbleSort(testCase.unSortedList)

			assertList(t, i, got, testCase.expectedSortedList)
			assertPass(t, i, pass, testCase.expetedPass)
		}
	})
}

func assertList(t testing.TB, i, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("[#%d] got: %v, but want: %v", i, got, want)
	}
}

func assertPass(t testing.TB, i, got, want int) {
	t.Helper()
	if !(got == want) {
		t.Errorf("[#%d] actual pass: %d, but expected pass: %d", i, got, want)
	}
}

package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("test to sort numbers", func(t *testing.T) {
		testCases := []struct {
			unSortedList       []int
			expectedSortedList []int
			expetedPass        int
			err                error
		}{
			{unSortedList: []int{}, expectedSortedList: []int{}, expetedPass: 0, err: fmt.Errorf("empty slice/array can't sorted")},
			{unSortedList: []int{1, 2}, expectedSortedList: []int{1, 2}, expetedPass: 1},
			{unSortedList: []int{1, 3, 2}, expectedSortedList: []int{1, 2, 3}, expetedPass: 2},
			{unSortedList: []int{10, 1, 3, 2, 5, 9, 4, 7, 8, 6}, expectedSortedList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expetedPass: 5},
		}

		for i, testCase := range testCases {
			got, pass, err := BubbleSort(testCase.unSortedList)

			assertError(t, i, err, testCase.err)
			assertList(t, i, got, testCase.expectedSortedList)
			assertPass(t, i, pass, testCase.expetedPass)
		}
	})

	t.Run("test to sort alphabets", func(t *testing.T) {
		testCases := []struct {
			unSortedList       []string
			expectedSortedList []string
			expetedPass        int
			err                error
		}{
			{unSortedList: []string{}, expectedSortedList: []string{}, expetedPass: 0, err: fmt.Errorf("empty slice/array can't sorted")},
			{unSortedList: []string{"apple", "dog"}, expectedSortedList: []string{"apple", "dog"}, expetedPass: 1},
			{unSortedList: []string{"apple", "dog", "banana"}, expectedSortedList: []string{"apple", "banana", "dog"}, expetedPass: 2},
		}

		for i, testCase := range testCases {
			got, pass, err := BubbleSort(testCase.unSortedList)

			assertError(t, i, err, testCase.err)
			assertList(t, i, got, testCase.expectedSortedList)
			assertPass(t, i, pass, testCase.expetedPass)
		}
	})

	t.Run("test to list of unsupported type", func(t *testing.T) {
		testCases := []struct {
			given interface{}
			err   error
		}{
			{given: true, err: fmt.Errorf(`only supported an array/slice type. But entered "bool" type`)},
			{given: []bool{true, false}, err: fmt.Errorf(`unsupported type "bool"`)},
			{given: []byte("hello"), err: fmt.Errorf(`unsupported type "uint8"`)},
		}

		for i, testCase := range testCases {
			got, _, err := BubbleSort(testCase.given)

			assertError(t, i, err, testCase.err)
			assertList(t, i, got, testCase.given)
		}
	})
}

func assertError(t testing.TB, i int, got, want error) {
	t.Helper()

	if got == nil && want == nil {
		t.Skip()
	}

	if (got == nil && want != nil) || (got != nil && want == nil) {
		t.Errorf("[#%d] got error: %v, but want: %v", i, got, want)
		t.FailNow()
	}

	if got.Error() != want.Error() {
		t.Errorf("[#%d] got error context: %q, but want: %q", i, got.Error(), want.Error())
	}
}

func assertList(t testing.TB, i int, got, want interface{}) {
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

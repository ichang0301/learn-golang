package linear_search_test

import (
	"testing"

	linear_search "github.com/ichang0301/learn-golang/32_search"
)

func TestLinearSearchUsingNumList(t *testing.T) {
	testcases := []struct {
		list        []int
		targetNum   int
		expectedPnt int
		isError     bool
	}{
		{
			list:        []int{7, 10, 3, 6, 4, 1, 2, 8, 9, 5},
			targetNum:   4,
			expectedPnt: 4,
			isError:     false,
		},
		{
			list:        []int{7, 10, 3, 6, 4, 1, 2, 8, 9, 5},
			targetNum:   11,
			expectedPnt: 0,
			isError:     true,
		},
		{
			list:        []int{-5, 1, 2, -3, 0, 4, 3, -4, -2, -1},
			targetNum:   -4,
			expectedPnt: 7,
			isError:     false,
		},
	}

	for _, tt := range testcases {
		// given
		linearSearch := linear_search.NewLinearSearch(tt.list)

		// when
		actualPnt, err := linearSearch.Search(tt.targetNum)

		// then
		if tt.isError && err == nil {
			t.Fatalf("expected error occured but didn't")
		} else if !tt.isError && err != nil {
			t.Fatalf("unexpected error occured: %#v", err)
		}

		if tt.expectedPnt != actualPnt {
			t.Errorf("expected point: %d, bug actual: %d", tt.expectedPnt, actualPnt)
		}
	}
}

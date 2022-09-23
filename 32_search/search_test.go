package search_test

import (
	"fmt"
	"testing"

	"github.com/ichang0301/learn-golang/32_search/linear_search"
	"github.com/ichang0301/learn-golang/utils"
	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

func TestSearchInt(t *testing.T) {
	testcases := []struct {
		description string
		list        []int
		target      int
		expectedPnt int
		expectedErr error
	}{
		{
			description: "no_such_element",
			list:        []int{7, 10, 3, 6, 4, 1, 2, 8, 9, 5},
			target:      11,
			expectedPnt: -1,
			expectedErr: errors.ErrNoSuchElement,
		},
		{
			description: "success_case",
			list:        []int{7, 10, 3, 6, 4, 1, 2, 8, 9, 5},
			target:      4,
			expectedPnt: 4,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// given
			linearSearch := linear_search.NewLinearSearch(tt.list)

			// when
			actualPnt, err := linearSearch.Search(tt.target)

			// then
			utils.AssertError(t, tt.expectedErr, err)
			if tt.expectedPnt != actualPnt {
				t.Errorf("expected point: %d, bug actual: %d", tt.expectedPnt, actualPnt)
			}
		})
	}
}

func TestSearchFloat(t *testing.T) {
	testcases := []struct {
		description string
		list        []float64
		target      float64
		expectedPnt int
		expectedErr error
	}{
		{
			description: "no_such_element",
			list:        []float64{0.1, 0.2, 1},
			target:      0,
			expectedPnt: -1,
			expectedErr: errors.ErrNoSuchElement,
		},
		{
			description: "success_case",
			list:        []float64{0.1, 0.2, 1},
			target:      1,
			expectedPnt: 2,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// given
			linearSearch := linear_search.NewLinearSearch(tt.list)

			// when
			actualPnt, err := linearSearch.Search(tt.target)

			// then
			utils.AssertError(t, tt.expectedErr, err)
			if tt.expectedPnt != actualPnt {
				t.Errorf("expected point: %d, bug actual: %d", tt.expectedPnt, actualPnt)
			}
		})
	}
}

func TestSearchString(t *testing.T) {
	testcases := []struct {
		description string
		list        []string
		target      string
		expectedPnt int
		expectedErr error
	}{
		{
			description: "no_such_element",
			list:        []string{"hello", "world", "golang"},
			target:      "invalid",
			expectedPnt: -1,
			expectedErr: errors.ErrNoSuchElement,
		},
		{
			description: "success_case",
			list:        []string{"hello", "world", "golang"},
			target:      "hello",
			expectedPnt: 0,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// given
			linearSearch := linear_search.NewLinearSearch(tt.list)

			// when
			actualPnt, err := linearSearch.Search(tt.target)

			// then
			utils.AssertError(t, tt.expectedErr, err)
			if tt.expectedPnt != actualPnt {
				t.Errorf("expected point: %d, bug actual: %d", tt.expectedPnt, actualPnt)
			}
		})
	}
}

func TestUnsupportedDataType(t *testing.T) {
	testcases := []struct {
		description string
		list        interface{}
		target      interface{}
	}{
		{
			description: "not_list",
			list:        "not_list",
			target:      "not_list",
		},
		{
			description: "slice_of_bool",
			list:        []bool{true, false},
			target:      true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			kind, _ := list.DetectDataType(tt.list)
			fmt.Printf("kind: %v\n", kind)

			// given
			linearSearch := linear_search.NewLinearSearch(tt.list)

			// when
			_, err := linearSearch.Search(true)

			// then
			utils.AssertError(t, errors.ErrUnsupportedType, err)
		})
	}
}

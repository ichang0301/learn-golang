package list_test

import (
	"reflect"
	"testing"

	"github.com/ichang0301/learn-golang/utils"
	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

func TestDetectDataType(t *testing.T) {
	testcases := []struct {
		description  string
		list         interface{}
		expectedType reflect.Kind
		expectedErr  error
	}{
		{
			description:  "not_list",
			list:         "not list",
			expectedType: reflect.Invalid,
			expectedErr:  errors.ErrUnsupportedType,
		},
		{
			description:  "empty_list",
			list:         []int{},
			expectedType: reflect.Invalid,
			expectedErr:  errors.ErrEmptyList,
		},
		{
			description:  "slice_of_int",
			list:         []int{1, 2, 3},
			expectedType: reflect.Int,
		},
		{
			description:  "slice_of_string",
			list:         []string{"a", "b", "c"},
			expectedType: reflect.String,
		},
		{
			description:  "slice_of_bool",
			list:         []bool{true, false},
			expectedType: reflect.Bool,
		},
		{
			description:  "slice_of_float",
			list:         []float64{0.1, 0.2, 0.3},
			expectedType: reflect.Float64,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// when
			actual, err := list.DetectDataType(tt.list)

			// then
			utils.AssertError(t, tt.expectedErr, err)
			if tt.expectedType != actual {
				t.Errorf("got: %v, but want: %v", actual, tt.expectedType)
			}
		})
	}
}

func TestGetIntSliceFromInterface(t *testing.T) {
	testcases := []struct {
		description  string
		expectedList interface{}
		expectedErr  error
	}{
		{
			description:  "not_slice",
			expectedList: 1,
			expectedErr:  errors.ErrUnsupportedType,
		},
		{
			description:  "not_int_slice",
			expectedList: []string{"a"},
			expectedErr:  errors.ErrUnsupportedType,
		},
		{
			description:  "empty_slice",
			expectedList: []int{},
			expectedErr:  errors.ErrEmptyList,
		},
		{
			description:  "slice_of_int",
			expectedList: []int{1, 2, 3},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// when
			actual, err := list.GetIntSliceFromInterface(tt.expectedList)

			// then
			utils.AssertError(t, tt.expectedErr, err)
			if err == nil {
				utils.AssertList(t, tt.expectedList, actual)
			}
		})
	}
}

func TestGetStringSliceFromInterface(t *testing.T) {
	testcases := []struct {
		description  string
		expectedList interface{}
		expectedErr  error
	}{
		{
			description:  "not_slice",
			expectedList: "abc",
			expectedErr:  errors.ErrUnsupportedType,
		},
		{
			description:  "not_string_slice",
			expectedList: []int{1},
			expectedErr:  errors.ErrUnsupportedType,
		},
		{
			description:  "empty_slice",
			expectedList: []string{},
			expectedErr:  errors.ErrEmptyList,
		},
		{
			description:  "slice_of_string",
			expectedList: []string{"a", "b"},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// when
			actual, err := list.GetStringSliceFromInterface(tt.expectedList)

			// then
			utils.AssertError(t, tt.expectedErr, err)
			if err == nil {
				utils.AssertList(t, tt.expectedList, actual)
			}
		})
	}
}

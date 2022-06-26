// http://www.eecs.qmul.ac.uk/~mmh/DCS128/2006/resources/selsort.html

package selection_sort

import (
	"math"
	"reflect"
	"strings"

	sort "github.com/ichang0301/learn-golang/31_sort"
)

type selectionSortAlgorithm struct {
	list interface{}
}

// NewSelectionSortAlgorithm creates a SelectionSortAlgorithm implemented SortAlgorithm
func NewSelectionSortAlgorithm(list interface{}) (*selectionSortAlgorithm, error) {
	s := selectionSortAlgorithm{list: list}
	if err := s.sort(); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *selectionSortAlgorithm) GetList() interface{} {
	return s.list
}

func (s *selectionSortAlgorithm) sort() error { // TODO: change the start letter to small letter
	kind, err := sort.DetectType(s.list)
	if err != nil {
		return err
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList := sort.InterfaceToSliceOfInt(s.list)
		s.list = SelectionSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList := sort.InterfaceToSliceOfString(s.list)
		s.list = SelectionSortStrings(stringList)

	// unsupported type
	default:
		return sort.ErrUnsupportedType
	}

	return nil
}

// SelectionSortNumbers sorts a list of numbers that the type is int using selection sort algorithm
func SelectionSortNumbers(list []int) []int {
	for i := 0; i < len(list); i++ {
		smallest := math.MaxInt
		index := i
		for j := i; j < len(list); j++ {
			if list[j] < smallest {
				smallest = list[j]
				index = j
			}
		}

		if index != i {
			list[i], list[index] = list[index], list[i]
		}
	}

	return list
}

// SelectionSortNumbers sorts a list of strings using selection sort algorithm
func SelectionSortStrings(list []string) []string {
	for i := len(list) - 1; i >= 0; i-- {
		biggestStr := "a"
		index := i
		for j := i; j >= 0; j-- {
			if strings.Compare(list[j], biggestStr) > 0 {
				biggestStr = list[j]
				index = j
			}
		}

		if index != i {
			list[i], list[index] = list[index], list[i]
		}
	}

	return list
}

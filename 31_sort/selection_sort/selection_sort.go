// http://www.eecs.qmul.ac.uk/~mmh/DCS128/2006/resources/selsort.html

package selection_sort

import (
	"math"
	"reflect"
	"strings"

	sort "github.com/ichang0301/learn-golang/31_sort"
)

type SelectionSortAlgorithm struct {
	list interface{}
}

func NewSelectionSortAlgorithm(list interface{}) SelectionSortAlgorithm {
	return SelectionSortAlgorithm{list: list}
}

func (s SelectionSortAlgorithm) GetSortedList() (sortedList interface{}, err error) {
	return s.Sort(s.list)
}

func (s SelectionSortAlgorithm) Sort(unSortedList interface{}) (sortedList interface{}, err error) {
	sortedList = unSortedList
	kind, err := sort.DetectType(unSortedList)
	if err != nil {
		return
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList := sort.InterfaceToSliceOfInt(unSortedList)
		sortedList = SelectionSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList := sort.InterfaceToSliceOfString(unSortedList)
		sortedList = SelectionSortStrings(stringList)
	// unsupported type
	default:
		err = sort.ErrUnsupportedType
	}

	return
}

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
			tmp := list[i]
			list[i] = list[index]
			list[index] = tmp
		}
	}

	return list
}

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
			tmp := list[i]
			list[i] = list[index]
			list[index] = tmp
		}
	}

	return list
}

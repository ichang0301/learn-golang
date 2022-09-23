// https://www.bbc.co.uk/bitesize/guides/zjdkw6f/revision/5

package merge_sort

import (
	"reflect"

	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

type mergeSortAlgorithm struct {
	list interface{}
}

func NewMergeSortAlgorithm(list interface{}) (*mergeSortAlgorithm, error) {
	m := mergeSortAlgorithm{list: list}
	if err := m.sort(); err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *mergeSortAlgorithm) GetList() interface{} {
	return m.list
}

func (m *mergeSortAlgorithm) sort() error {
	kind, err := list.DetectWhatListDataTypeIs(m.list)
	if err != nil {
		return err
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList, err := list.GetIntSliceFromInterface(m.list)
		if err != nil {
			return err
		}
		m.list = MergeSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList, err := list.GetStringSliceFromInterface(m.list)
		if err != nil {
			return err
		}
		m.list = MergeSortStrings(stringList)

	// unsupported type
	default:
		return errors.ErrUnsupportedType
	}

	return nil
}

func MergeSortNumbers(list []int) []int {
	if len(list) < 2 {
		return list
	}
	mid := (len(list)) / 2
	return MergeNumbers(MergeSortNumbers(list[:mid]), MergeSortNumbers(list[mid:]))
}

// MergeNumbers left and right slice into newly created slice: https://austingwalters.com/merge-sort-in-go-golang/
func MergeNumbers(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func MergeSortStrings(list []string) []string {
	if len(list) < 2 {
		return list
	}
	mid := (len(list)) / 2
	return MergeStrings(MergeSortStrings(list[:mid]), MergeSortStrings(list[mid:]))
}

// MergeStrings left and right slice into newly created slice: https://austingwalters.com/merge-sort-in-go-golang/
func MergeStrings(left, right []string) []string {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]string, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

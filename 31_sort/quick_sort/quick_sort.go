// http://syllabus.cs.manchester.ac.uk/ugt/2019/COMP26120/SortingTool/quick_sort_info.html

package quick_sort

import (
	"reflect"
	"strings"

	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

type quickSortAlgorithm struct {
	list interface{}
}

func NewQuickSortAlgorithm(list interface{}) (*quickSortAlgorithm, error) {
	i := quickSortAlgorithm{list: list}
	if err := i.sort(); err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *quickSortAlgorithm) GetList() interface{} {
	return i.list
}

func (i *quickSortAlgorithm) sort() error {
	kind, err := list.DetectWhatListDataTypeIs(i.list)
	if err != nil {
		return err
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList, err := list.GetIntSliceFromInterface(i.list)
		if err != nil {
			return err
		}
		i.list = QuickSortNumbers(intList, 0, len(intList)-1)

	// sort the list of string type
	case reflect.String:
		stringList, err := list.GetStringSliceFromInterface(i.list)
		if err != nil {
			return err
		}
		i.list = QuickSortStrings(stringList, 0, len(stringList)-1)

	// unsupported type
	default:
		return errors.ErrUnsupportedType
	}

	return nil
}

func QuickSortNumbers(list []int, low, high int) []int {
	if low < high {
		var p int
		list, p = partitionNumbers(list, low, high)
		list = QuickSortNumbers(list, low, p-1)
		list = QuickSortNumbers(list, p+1, high)
	}
	return list
}

func partitionNumbers(list []int, low, high int) ([]int, int) {
	pivot := list[high]
	i := low
	for j := low; j < high; j++ {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	list[i], list[high] = list[high], list[i]
	return list, i
}

func QuickSortStrings(list []string, low, high int) []string {
	if low < high {
		var p int
		list, p = partitionStrings(list, low, high)
		list = QuickSortStrings(list, low, p-1)
		list = QuickSortStrings(list, p+1, high)
	}
	return list
}

func partitionStrings(list []string, low, high int) ([]string, int) {
	pivot := list[high]
	i := low
	for j := low; j < high; j++ {
		if strings.Compare(list[j], pivot) < 0 {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	list[i], list[high] = list[high], list[i]
	return list, i
}

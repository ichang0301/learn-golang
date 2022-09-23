// https://corewar.co.uk/assembly/insertion.htm

package insertion_sort

import (
	"reflect"
	"strings"

	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

type insertionSortAlgorithm struct {
	list interface{}
}

func NewInsertionSortAlgorithm(list interface{}) (*insertionSortAlgorithm, error) {
	i := insertionSortAlgorithm{list: list}
	if err := i.sort(); err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *insertionSortAlgorithm) GetList() interface{} {
	return i.list
}

func (i *insertionSortAlgorithm) sort() error {
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
		i.list = InsertionSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList, err := list.GetStringSliceFromInterface(i.list)
		if err != nil {
			return err
		}
		i.list = InsertionSortStrings(stringList)

	// unsupported type
	default:
		return errors.ErrUnsupportedType
	}

	return nil
}

func InsertionSortNumbers(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}

	return list
}

func InsertionSortStrings(list []string) []string {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if strings.Compare(list[j], list[j-1]) < 0 {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}

	return list
}

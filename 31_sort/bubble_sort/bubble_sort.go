// https://www.bbc.co.uk/bitesize/guides/z2m3b9q/revision/2

// The bubble sort algorithm has a worst-case time complexity of O(n2). The bubble sort has a space complexity of O(1).

package bubble_sort

import (
	"reflect"
	"strings"

	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

type bubbleSortAlgorithm struct {
	list interface{}
	pass int
}

func NewBubbleSortAlgorithm(list interface{}) (*bubbleSortAlgorithm, error) {
	b := bubbleSortAlgorithm{list: list}
	if err := b.sort(); err != nil {
		return nil, err
	}
	return &b, nil
}

func (b *bubbleSortAlgorithm) GetList() interface{} {
	return b.list
}

func (b *bubbleSortAlgorithm) GetPass() int {
	return b.pass
}

func (b *bubbleSortAlgorithm) sort() error {
	kind, err := list.DetectDataType(b.list)
	if err != nil {
		return err
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList, err := list.GetIntSliceFromInterface(b.list)
		if err != nil {
			return err
		}
		b.list, b.pass = BubbleSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList, err := list.GetStringSliceFromInterface(b.list)
		if err != nil {
			return err
		}
		b.list, b.pass = BubbleSortStrings(stringList)

	// unsupported type
	default:
		return errors.ErrUnsupportedType
	}

	return nil
}

// BubbleSortNumbers sorts a list of numbers that the type is int using bubble sort algorithm
func BubbleSortNumbers(list []int) ([]int, int) {
	var isSorted bool
	var pass int
	for i := pass; i < len(list); i++ {
		if !isSorted {
			isSorted = true
			for j := 0; j < len(list)-1; j++ {
				if list[j] > list[j+1] {
					list[j+1], list[j] = list[j], list[j+1]
					isSorted = false
				}
			}
		} else {
			break
		}
		pass++
	}

	return list, pass
}

// BubbleSortStrings sorts a list of strings using bubble sort algorithm
func BubbleSortStrings(list []string) ([]string, int) {
	var isSorted bool
	var pass int
	for i := pass; i < len(list); i++ {
		if !isSorted {
			isSorted = true
			for j := 0; j < len(list)-1; j++ {
				if strings.Compare(list[j], list[j+1]) > 0 {
					list[j+1], list[j] = list[j], list[j+1]
					isSorted = false
				}
			}
		} else {
			break
		}
		pass++
	}

	return list, pass
}

// https://www.bbc.co.uk/bitesize/guides/z2m3b9q/revision/2

// The bubble sort algorithm has a worst-case time complexity of O(n2). The bubble sort has a space complexity of O(1).

package bubble_sort

import (
	"reflect"
	"strings"

	sort "github.com/ichang0301/learn-golang/31_sort"
)

type BubbleSortAlgorithm struct{}

func (b BubbleSortAlgorithm) Sort(unSortedList interface{}) (sortedList interface{}, err error) {
	sortedList, _, err = BubbleSort(unSortedList)
	return
}

// BubbleSort sorts items using the bubble sort algorithm and then returns the sorted list and how many times the algorithm goes through the list.
func BubbleSort(list interface{}) (result interface{}, pass int, err error) { // pass: how many times we take of sorting the list using bubble sort algorithm
	result = list
	kind, err := sort.DetectType(list)
	if err != nil {
		return
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList := sort.InterfaceToSliceOfInt(list)
		result, pass = BubbleSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList := sort.InterfaceToSliceOfString(list)
		result, pass = BubbleSortStrings(stringList)

	// unsupported type
	default:
		err = sort.ErrUnsupportedType
	}
	return
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
					tmp := list[j+1]
					list[j+1] = list[j]
					list[j] = tmp
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
					tmp := list[j+1]
					list[j+1] = list[j]
					list[j] = tmp
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

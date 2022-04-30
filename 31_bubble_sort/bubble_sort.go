// https://www.bbc.co.uk/bitesize/guides/z2m3b9q/revision/2

// The bubble sort algorithm has a worst-case time complexity of O(n2). The bubble sort has a space complexity of O(1).

package bubble_sort

import (
	"reflect"
)

// BubbleSort sorts items using the bubble sort algorithm and then returns the sorted list and how many times the algorithm goes through the list.
func BubbleSort(list interface{}) (result interface{}, pass int) { // pass: how many times we take of sorting the list using bubble sort algorithm
	var isSorted bool
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array:
		l := reflect.ValueOf(list)
		result = l.Interface()

		length := l.Len()
		if length == 0 {
			return
		}

		switch l.Index(0).Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var sortedList []int

			for i := 0; i < length; i++ {
				sortedList = append(sortedList, int(l.Index(i).Int()))
			}

			for i := pass; i < length; i++ {
				if !isSorted {
					isSorted = true
					for j := pass; j < length-1; j++ {
						if sortedList[j] > sortedList[j+1] {
							tmp := sortedList[j+1]
							sortedList[j+1] = sortedList[j]
							sortedList[j] = tmp
							isSorted = false
						}
					}
				} else {
					break
				}
				pass++
			}

			result = sortedList
		}

	}

	return
}
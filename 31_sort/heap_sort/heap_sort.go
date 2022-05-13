// http://syllabus.cs.manchester.ac.uk/ugt/2019/COMP26120/SortingTool/heap_sort_info.html

package heap_sort

import (
	"container/heap"
	"reflect"
	"strings"

	sort "github.com/ichang0301/learn-golang/31_sort"
)

type heapSortAlgorithm struct {
	list interface{}
}

func NewHeapSortAlgorithm(list interface{}) (*heapSortAlgorithm, error) {
	i := heapSortAlgorithm{list: list}
	if err := i.sort(); err != nil {
		return nil, err
	}
	return &i, nil
}

func (i *heapSortAlgorithm) GetList() interface{} {
	return i.list
}

func (i *heapSortAlgorithm) sort() error {
	kind, err := sort.DetectType(i.list)
	if err != nil {
		return err
	}

	switch kind {
	// sort the list of integer type
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intList := sort.InterfaceToSliceOfInt(i.list)
		i.list = HeapSortNumbers(intList)

	// sort the list of string type
	case reflect.String:
		stringList := sort.InterfaceToSliceOfString(i.list)
		i.list = HeapSortStrings(stringList)

	// unsupported type
	default:
		return sort.ErrUnsupportedType
	}

	return nil
}

// HeapSortNumbers sorts using "container/heap" package
func HeapSortNumbers(list []int) []int {
	h := intHeap(list)
	heap.Init(&h)

	var result []int
	for h.Len() > 0 {
		result = append(result, heap.Pop(&h).(int))
	}

	return result
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Pop() (n interface{}) {
	old := *h
	l := len(old)
	n = old[l-1]
	*h = old[0 : l-1]
	return
}

func (h *intHeap) Push(n interface{}) {}

func HeapSortStrings(list []string) []string {
	h := StrHeap(list)
	heap.Init(&h)

	var result []string
	for h.Len() > 0 {
		result = append(result, heap.Pop(&h).(string))
	}

	return result
}

// StrHeap is a min-heap for slice of string
type StrHeap []string

func (h StrHeap) Len() int {
	return len(h)
}

func (h StrHeap) Less(i, j int) bool {
	return strings.Compare(h[i], h[j]) < 0
}

func (h StrHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StrHeap) Pop() (s interface{}) {
	old := *h
	l := len(old)
	s = old[l-1]
	*h = old[0 : l-1]
	return
}

func (h *StrHeap) Push(s interface{}) {}

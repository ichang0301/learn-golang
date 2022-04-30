package sort

type SortAlgorithm interface {
	Sort(unSortedList interface{}) (sortedList interface{}, err error)
}

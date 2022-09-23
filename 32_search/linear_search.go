package linear_search

import "fmt"

type linearSearch struct {
	list []int
}

func NewLinearSearch(list []int) *linearSearch {
	return &linearSearch{list: list}
}

func (l *linearSearch) Search(n int) (int, error) {
	for i, x := range l.list {
		if x == n {
			return i, nil
		}
	}

	return 0, fmt.Errorf("there is no %d in the list", n)
}

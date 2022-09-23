package linear_search

import (
	"reflect"

	"github.com/ichang0301/learn-golang/utils/errors"
	"github.com/ichang0301/learn-golang/utils/list"
)

type linearSearch struct {
	list interface{}
}

func NewLinearSearch(list interface{}) *linearSearch {
	return &linearSearch{list: list}
}

func (l *linearSearch) Search(input interface{}) (int, error) {
	kind, err := list.DetectWhatListDataTypeIs(l.list)
	if err != nil {
		return 0, err
	}

	if kind != reflect.TypeOf(input).Kind() {
		return 0, errors.ErrNoSuchElement
	}

	switch kind {
	// search integer in the list
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return l.linearSearchInt(input.(int))

	// search float in the list
	case reflect.Float32, reflect.Float64:
		return l.linearSearchFloat(input.(float64))

	// search string in the list
	case reflect.String:
		return l.linearSearchString(input.(string))

	// unsupported type
	default:
		return -1, errors.ErrUnsupportedType
	}
}

func (l *linearSearch) linearSearchInt(input int) (int, error) {
	intList, err := list.GetIntSliceFromInterface(l.list)
	if err != nil {
		return 0, err
	}

	for i, x := range intList {
		if x == input {
			return i, nil
		}
	}

	return -1, errors.ErrNoSuchElement
}

func (l *linearSearch) linearSearchFloat(input float64) (int, error) {
	floatList, err := list.GetFloatSliceFromInterface(l.list)
	if err != nil {
		return 0, err
	}

	for i, x := range floatList {
		if x == input {
			return i, nil
		}
	}

	return -1, errors.ErrNoSuchElement
}

func (l *linearSearch) linearSearchString(input string) (int, error) {
	stringList, err := list.GetStringSliceFromInterface(l.list)
	if err != nil {
		return 0, err
	}

	for i, x := range stringList {
		if x == input {
			return i, nil
		}
	}

	return -1, errors.ErrNoSuchElement
}

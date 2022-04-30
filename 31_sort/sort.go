package sort

import "reflect"

const (
	ErrUnsupportedType = SortErrorFormat("unsupported type")
	ErrEmptyList       = SortErrorFormat("empty slice/array can't sorted")
)

// SortErrorFormat is the format that prints errors occurred when sorting a list
type SortErrorFormat string

// Error returns a sort error context for implementing of error
func (e SortErrorFormat) Error() string { // implement 'error' interface. : https://go.dev/blog/error-handling-and-go
	return string(e)
}

// SortAlgorithm is an interface to sort a list
type SortAlgorithm interface {
	Sort(unSortedList interface{}) (sortedList interface{}, err error)
}

func DetectType(input interface{}) (reflect.Kind, error) {
	v := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice, reflect.Array:
		length := v.Len()
		if length == 0 {
			return 0, ErrEmptyList
		}
		return v.Index(0).Type().Kind(), nil

	// only supported array/slice type
	default:
		return 0, ErrUnsupportedType
	}
}

func InterfaceToSliceOfInt(input interface{}) []int {
	var result []int
	v := reflect.ValueOf(input)
	for i := 0; i < v.Len(); i++ {
		result = append(result, int(v.Index(i).Int()))
	}
	return result
}

func InterfaceToSliceOfString(input interface{}) []string {
	var result []string
	v := reflect.ValueOf(input)
	for i := 0; i < v.Len(); i++ {
		result = append(result, v.Index(i).String())
	}
	return result
}

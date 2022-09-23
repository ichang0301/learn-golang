package list

import (
	"reflect"

	"github.com/ichang0301/learn-golang/utils/errors"
)

// DetectDataType gets a data type from the interface type.
// But it supports only a slice or array.
func DetectDataType(input interface{}) (reflect.Kind, error) {
	v := reflect.ValueOf(input)

	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice, reflect.Array:
		length := v.Len()
		if length == 0 {
			return 0, errors.ErrEmptyList
		}
		return v.Index(0).Type().Kind(), nil

	// only supported array/slice type
	default:
		return 0, errors.ErrUnsupportedType
	}
}

// GetIntSliceFromInterface gets a slice of int from the interface
func GetIntSliceFromInterface(input interface{}) ([]int, error) {
	typeName, err := DetectDataType(input)
	if err != nil {
		return nil, err
	}

	switch typeName {
	// return a number list
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return getIntSliceFromInterface(input), nil
	}

	// unsupported type
	return nil, errors.ErrUnsupportedType
}

func getIntSliceFromInterface(input interface{}) []int {
	var result []int
	v := reflect.ValueOf(input)
	for i := 0; i < v.Len(); i++ {
		result = append(result, int(v.Index(i).Int()))
	}
	return result
}

// GetStringSliceFromInterface gets a slice of string from the interface
func GetStringSliceFromInterface(input interface{}) ([]string, error) {
	typeName, err := DetectDataType(input)
	if err != nil {
		return nil, err
	}

	switch typeName {
	// return a string list
	case reflect.String:
		return getStringSliceFromInterface(input), nil
	}

	// unsupported type
	return nil, errors.ErrUnsupportedType
}

func getStringSliceFromInterface(input interface{}) []string {
	var result []string
	v := reflect.ValueOf(input)
	for i := 0; i < v.Len(); i++ {
		result = append(result, v.Index(i).String())
	}
	return result
}

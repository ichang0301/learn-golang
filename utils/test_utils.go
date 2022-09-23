package utils

import (
	"reflect"
	"testing"
)

func AssertError(t testing.TB, expected, actual error) {
	t.Helper()

	if (actual == nil && expected != nil) || (actual != nil && expected == nil) {
		t.Fatalf("got error: %#v, but want: %#v", actual, expected)
	}

	if !(actual == nil && expected == nil) && (actual.Error() != expected.Error()) {
		t.Errorf("got error context: %q, but want: %q", actual.Error(), expected.Error())
	}
}

func AssertElement(t testing.TB, expected, actual interface{}) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v, but want: %v", actual, expected)
	}
}

package test_utils

import (
	"reflect"
	"testing"
)

func AssertError(t testing.TB, i int, got, want error) {
	t.Helper()

	if (got == nil && want != nil) || (got != nil && want == nil) {
		t.Fatalf("[#%d] got error: %v, but want: %v", i, got, want)
	}

	if !(got == nil && want == nil) && (got.Error() != want.Error()) {
		t.Errorf("[#%d] got error context: %q, but want: %q", i, got.Error(), want.Error())
	}
}

func AssertList(t testing.TB, i int, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("[#%d] got: %v, but want: %v", i, got, want)
	}
}

func AssertPass(t testing.TB, i, got, want int) {
	t.Helper()

	if !(got == want) {
		t.Errorf("[#%d] actual pass: %d, but expected pass: %d", i, got, want)
	}
}

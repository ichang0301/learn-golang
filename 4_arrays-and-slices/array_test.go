package arrays

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumArray(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
		// Structs formatted with %v show field values in their default formats.
		// The %+v form shows the fields by name, while %#v formats the struct in
		// Go source format.
		// person := struct {
		// 	Name string
		// 	Age  int
		// }{"Kim", 22}
		// fmt.Printf("%v %+v %#v\n", person, person, person)
		// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
	}
}

func TestSumSlice(t *testing.T) {
	t.Run("collection of any size of numbers: size is 3", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlices(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size of numbers: size is 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlices(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{1, 2, 3, 4, 5}

	got := SumAll(numbers1, numbers2)
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) { // reflect.DeepEqual is not "type safe"
		t.Errorf("got %d want %d given, %v %v", got, want, numbers1, numbers2)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) { // A handy side-effect of this is this adds a little type-safety to our code.
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// example for documantation http://localhost:6060/pkg/github.com/ichang0301/learn-golang/2_integer/ (after run `godoc -http=localhost:6060` in terminal): https://go.dev/blog/examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6go
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	assertEqualInts(t, sum, expected)
}

func assertEqualInts(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

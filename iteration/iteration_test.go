package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCompare(t *testing.T) {
	compared := Compare("a", "a")
	expected := true

	if compared != expected {
		t.Errorf("expected %t but got %t", expected, compared) //%t is used for boolean
	}
}

// writing benchmarks in go , similair to tests and measures how long
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}

//to run bench mark, -> go test -bench="." for WINDOWS, -> go test -bench=. for UNIX

// write example
func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	//Output: aaaaa
}

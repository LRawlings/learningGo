package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expect %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 9)
	fmt.Println(repeat)
	// Output: bbbbbbbbb
}

func TestContains(t *testing.T) {
	contains := strings.Contains("cider", "der")
	expected := true

	if contains != expected {
		t.Errorf("expect %v but got %v", expected, contains)
	}
}

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

func TestCount(t *testing.T) {
	count := strings.Count("cheddar valley", "e")
	expected := 2

	if count != expected {
		t.Errorf("expect %v but got %v", expected, count)
	}
}

func TestCountBlank(t *testing.T) {
	count := strings.Count("cheddar valley", "")
	expected := 15

	if count != expected {
		t.Errorf("expect %v but got %v", expected, count)
	}
}

func TestHasPrefix(t *testing.T) {
	hasPrefix := hasPrefix("Farmhouse Cider", "Farmhouse ")
	expected := true

	if hasPrefix != expected {
		t.Errorf("expect %v but got %v", expected, hasPrefix)
	}
}

package mylibrary_test

import (
	mylibrary "project/pkg/mylibrary"
	"testing"
)

func TestGreet(t *testing.T) {
	want := "Hello World"
	if got := mylibrary.Greet("World"); got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
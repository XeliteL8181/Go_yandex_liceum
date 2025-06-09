package lesson5

import "testing"

func TestMultiply(t *testing.T) {
	got := Multiply(13, 4)
	expected := 52

	if got != expected {
		t.Errorf(`Mulptiply(13, 4)=%q, want %q`, got, expected)
	}
}
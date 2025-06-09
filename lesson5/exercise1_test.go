package lesson5

import "testing"

func TestSum(t *testing.T) {
	got := Sum(2, 3)
	expected := 5

	if got != expected {
		t.Fatalf(`Sum(2, 3) = %q, want %q`, got, expected)
	}
}
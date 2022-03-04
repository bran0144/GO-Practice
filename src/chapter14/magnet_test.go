package chapter14

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Error(errorString2(2, 1, got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	want := 8
	got := Larger(4, 8)
	if got != want {
		t.Error(errorString2(4, 8, got, want))
	}
}

func errorString2(a int, b int, got int, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}

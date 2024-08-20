package slice_operations

import (
	"slices"
	"testing"
)

func TestRepeatSlice(t *testing.T) {
	s := []int{1, 2}
	r := slices.Repeat(s, 3)
	expected := []int{1, 2, 1, 2, 1, 2}

	if !slices.Equal(r, expected) {
		t.Errorf("Expected %v, but got %v", expected, r)
	}
}

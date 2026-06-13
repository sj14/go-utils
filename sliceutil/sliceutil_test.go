package sliceutil

import (
	"slices"
	"testing"
)

func TestShuffle(t *testing.T) {
	in := []int8{1, 2, 3, 4, 5, 6}
	copy := slices.Clone(in)

	Shuffle(in)

	if slices.Equal(in, copy) {
		t.Fail()
	}
}

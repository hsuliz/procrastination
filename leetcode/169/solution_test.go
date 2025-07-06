package _169

import (
	"testing"
)

func TestSolution(t *testing.T) {
	got := majorityElement([]int{8, 8, 7, 7, 7})
	if got != 7 {
		t.Errorf("got: %v, want: %v", got, 7)
	}
}

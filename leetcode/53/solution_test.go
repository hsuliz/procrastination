package _53

import "testing"

func TestMaxSubArray(t *testing.T) {
	given := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	got := maxSubArray(given)
	if got != 6 {
		t.Errorf("got: %v, want: %v", got, 6)
	}
}

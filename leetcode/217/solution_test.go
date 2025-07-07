package _217

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	given := []int{1, 2, 3, 1}
	got := containsDuplicate(given)

	if got != true {
		t.Errorf("got: %v, want: %v", got, true)
	}
}

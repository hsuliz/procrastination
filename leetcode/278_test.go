package leetcode

import (
	"testing"
)

func init() {
	isBadVersion = func(version int) bool {
		return version >= 60
	}
}

func TestFirstBadVersion(t *testing.T) {
	given := 100
	want := 60
	got := firstBadVersion(given)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

package leetcode

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	testCases := []struct {
		name  string
		given [2]string
		want  bool
	}{
		{"different chars", [...]string{"a", "b"}, false},
		{"inefficient chars", [...]string{"aa", "ab"}, false},
		{"enough chars", [...]string{"aa", "aab"}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canConstruct(tc.given[0], tc.given[1])
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

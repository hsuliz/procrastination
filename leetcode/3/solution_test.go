package _3

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		given string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{"dvdf", 3},
		{"abba", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.given, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.given)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

package _409

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name  string
		given string
		want  int
	}{
		{"long palindrome", "abccccdd", 7},
		{"small palindrome", "a", 1},
		{"same palindrome", "ccc", 3},
		{"bananas", "bananas", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := longestPalindrome(tc.given)
			if got != tc.want {
				t.Errorf("want %v, got %v,", tc.want, got)
			}
		})
	}
}

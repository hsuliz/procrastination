package leetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var testCases = []struct {
		name  string
		given string
		want  bool
	}{
		{"valid", "A man, a plan, a canal: Panama", true},
		{"non valid", "race a car", false},
		{"empty string", " ", true},
		{"silly", "0P", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.given)
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

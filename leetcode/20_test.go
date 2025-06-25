package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	var testCases = []struct {
		name  string
		given string
		want  bool
	}{
		{"() should be true", "()", true},
		{"()[]{} ", "()[]{}", true},
		{"(] should be false", "(]", false},
		{"] should be false", "]", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ans := isValid(tc.given)
			if ans != tc.want {
				t.Errorf("want %v, got %v,", tc.want, ans)
			}
		})
	}
}

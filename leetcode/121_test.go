package leetcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var testCases = []struct {
		name  string
		given []int
		want  int
	}{
		{"profitable", []int{7, 1, 5, 3, 6, 4}, 5},
		{"no profitable", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxProfit(tc.given)
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

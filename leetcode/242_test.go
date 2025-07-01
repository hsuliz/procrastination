package leetcode

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	//fmt.Println(isAnagram("ggii", "eekk"))
	//fmt.Println(isAnagram("anagram", "nagaram"))
	var testCases = []struct {
		name   string
		givenS string
		givenT string
		want   bool
	}{
		{"true anagram", "anagram", "nagaram", true},
		{"false anagram", "ggii", "eekk", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ans := isAnagram(tc.givenS, tc.givenT)
			if ans != tc.want {
				t.Errorf("want %v, got %v,", tc.want, ans)
			}
		})
	}
}

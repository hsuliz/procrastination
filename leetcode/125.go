package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var res []rune
	for _, val := range s {
		if unicode.IsNumber(val) || unicode.IsLetter(val) {
			res = append(res, val)
		}
	}

	for i, j := 0, len(res)-1; i < len(res)/2; i, j = i+1, j-1 {
		if res[i] != res[j] {
			return false
		}
	}

	return true
}

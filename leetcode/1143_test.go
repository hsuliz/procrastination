package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	x := longestCommonSubsequence(
		"abcde",
		"ace",
	)
	fmt.Println(x)

	x = longestCommonSubsequence(
		"abc",
		"def",
	)
	fmt.Println(x)
}

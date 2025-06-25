package leetcode

import (
	"fmt"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	x := subarraySum(
		[]int{1, 2, 3},
		3,
	)
	fmt.Println(x)
}

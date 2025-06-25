package leetcode

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	x := groupAnagrams(
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
	)
	fmt.Println(x)
}

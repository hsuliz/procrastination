package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	given := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	res := search(given, 6)
	fmt.Println(res)
}

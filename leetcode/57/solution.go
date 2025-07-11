package _57

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		prev := res[len(res)-1]
		if prev[1] >= curr[0] {
			if curr[1] > prev[1] {
				prev[1] = curr[1]
			}
			res[len(res)-1] = prev
		} else {
			res = append(res, curr)
		}
	}
	return res
}

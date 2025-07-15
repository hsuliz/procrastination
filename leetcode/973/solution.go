package _973

import (
	"slices"
)

func kClosest(points [][]int, k int) [][]int {
	slices.SortFunc(points, func(p1, p2 []int) int {
		xp := p1[0]*p1[0] + p1[1]*p1[1]
		yp := p2[0]*p2[0] + p2[1]*p2[1]
		return xp - yp
	})

	return points[:k]
}

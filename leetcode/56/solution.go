package _57

func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				tmp := intervals[j]
				intervals[j] = intervals[j+1]
				intervals[j+1] = tmp
			}
		}
	}

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		prev := res[len(res)-1]
		curr := intervals[i]
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

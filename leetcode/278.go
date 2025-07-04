package leetcode

var isBadVersion func(int) bool

func firstBadVersion(n int) int {
	start, end := 0, n
	for start <= end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

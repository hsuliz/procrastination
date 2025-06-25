package leetcode

func climbStairs(n int) int {
	cache := []int{1, 2, 3}
	for i := 3; i < n; i++ {
		cache = append(cache, cache[i-1]+cache[i-2])
	}
	return cache[n-1]
}

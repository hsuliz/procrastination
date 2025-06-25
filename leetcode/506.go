package leetcode

func subarraySum(nums []int, k int) int {
	prefixMap := map[int]int{0: 1}
	counter := 0
	sum := 0

	for _, num := range nums {
		sum += num
		if prefix, exists := prefixMap[sum-k]; exists {
			counter += prefix
		}
		prefixMap[sum]++
	}
	return counter
}

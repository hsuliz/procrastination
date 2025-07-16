package _3

func lengthOfLongestSubstring(s string) int {
	lastSeenMap := map[uint8]int{}
	var res int
	var start int
	for i := 0; i < len(s); i++ {
		val := s[i]
		if lastSeen, ok := lastSeenMap[val]; ok && lastSeen >= start {
			start = lastSeen + 1
		}
		windowSize := i - start + 1
		lastSeenMap[val] = i
		if windowSize > res {
			res = windowSize
		}
	}
	return res
}

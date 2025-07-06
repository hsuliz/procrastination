package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	tracker := map[rune]int{}

	for _, val := range magazine {
		tracker[val]++
	}

	for _, val := range ransomNote {
		if _, ok := tracker[val]; !ok {
			return false
		}
		tracker[val]--
		if tracker[val] == -1 {
			return false
		}
	}

	return true
}

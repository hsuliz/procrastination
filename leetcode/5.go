package leetcode

func longestPalindrome(s string) string {
	left, right := 0, len(s)-1
	result := ""
	for range len(s) {
		if s[left] == s[right] {
			left++
			right--
		}
	}
	return result
}

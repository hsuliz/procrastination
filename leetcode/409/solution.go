package _409

func longestPalindrome(s string) int {
	counter := map[rune]int{}
	for _, val := range s {
		counter[val]++
	}
	res := 0
	flag := false
	for _, val := range counter {
		res += (val / 2) * 2
		if val%2 == 1 {
			flag = true
		}
	}

	if flag {
		res++
	}

	return res
}

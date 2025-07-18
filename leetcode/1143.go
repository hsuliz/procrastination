package leetcode

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				cache[i+1][j+1] = cache[i][j] + 1
			} else {
				cache[i+1][j+1] = max(cache[i+1][j], cache[i][j+1])
			}
		}
	}

	fmt.Println(cache)
	return cache[m][n]
}

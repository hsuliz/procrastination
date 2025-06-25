package leetcode

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	anagMap := map[string][]string{}

	for _, word := range strs {
		charArr := make([]int, 26)
		for _, char := range word {
			charArr[char-'a']++
		}
		key := ""
		for _, char := range charArr {
			key += fmt.Sprintf("%d#", char)
		}
		anagMap[key] = append(anagMap[key], word)
	}

	var result [][]string
	for _, value := range anagMap {
		result = append(result, value)
	}
	return result
}

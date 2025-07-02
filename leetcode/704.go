package leetcode

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for {
		if high < low {
			return -1
		}
		mid := low + (high-low)/2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

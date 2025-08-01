package _15

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := num + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{num, nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return res
}

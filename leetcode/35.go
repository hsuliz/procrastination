package leetcode

func searchInsert(nums []int, target int) int {
	first, last := 0, len(nums)-1
	for {
		if first > last {
			return first
		}

		mid := first + (last-first)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			last = mid - 1
			continue
		} else {
			first = mid + 1
			continue
		}
	}
}

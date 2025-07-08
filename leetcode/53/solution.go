package _53

func maxSubArray(nums []int) int {
	maxEnd, res := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+maxEnd > nums[i] {
			maxEnd += nums[i]
		} else {
			maxEnd = nums[i]
		}

		if res < maxEnd {
			res = maxEnd
		}
	}

	return res
}

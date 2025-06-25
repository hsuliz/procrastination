package leetcode

func twoSum(nums []int, target int) []int {
	reqMap := map[int]int{}

	for i, num := range nums {
		need := target - num
		if req, ok := reqMap[need]; ok {
			return []int{req, i}
		}
		reqMap[num] = i
	}

	return nil
}

func twoSumN2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

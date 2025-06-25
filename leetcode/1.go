package leetcode

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for i, num := range nums {
		toFind := target - num
		if toFindIndex, exists := numMap[toFind]; exists {
			return []int{
				toFindIndex, i,
			}
		}
		numMap[num] = i
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

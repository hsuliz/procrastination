package leetcode

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsSet := map[int]bool{}
	for _, num := range nums {
		numsSet[num] = true
	}

	var maxSequence = 0
	for num, _ := range numsSet {
		if !numsSet[num-1] {
			currentNum := num
			sequence := 1
			for numsSet[currentNum+1] {
				currentNum++
				sequence++
			}

			if sequence > maxSequence {
				maxSequence = sequence
			}
			if maxSequence > len(nums)/2 {
				break
			}
		}
	}

	return maxSequence
}

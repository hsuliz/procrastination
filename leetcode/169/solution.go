package _169

func majorityElement(nums []int) int {
	var count, candidate int
	for _, num := range nums {
		switch {
		case count == 0:
			candidate = num
			fallthrough
		case candidate == num:
			count++
		default:
			count--
		}
	}
	return candidate
}

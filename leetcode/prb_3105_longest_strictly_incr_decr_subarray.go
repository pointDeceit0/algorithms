package leetcode


func longestMonotonicSubarray(nums []int) int {
	prev := nums[0]
	cntIncr, cntDecr := 1, 1
	maxCnt := 1
	for _, s := range nums[1:] {
		if s < prev {
			cntDecr++
			cntIncr = 1
			maxCnt = max(maxCnt, cntDecr)
		} else if s > prev {
			cntIncr++
			cntDecr = 1
			maxCnt = max(maxCnt, cntIncr)
		} else {
			cntIncr, cntDecr = 1, 1
		}
		prev = s
	}

	return max(maxCnt, cntIncr, cntDecr)
}
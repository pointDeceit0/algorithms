package leetcode

import (
	"slices"
)

func firstStableIndex(nums []int, k int) int {

	min := nums[len(nums) - 1]
	mins := make([]int, len(nums))
	for i, v := range slices.Backward(nums) {
		if v < min {
			mins[i] = v
			min = v
		} else {
			mins[i] = min
		}
	}

	max := nums[0]
	ans := -1
	for i, v := range nums {
		if v > max {
			max = v
		}
		if max - mins[i] <= k {
			ans = i
			break
		}
	}

	return ans
}
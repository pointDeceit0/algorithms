package leetcode

import (
	"slices"
)


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func insertBinary(a *[]int, v int) {
	idx, _ := slices.BinarySearch(*a, v)
	*a = slices.Insert(*a, idx, v)
}


func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	// TLE

	nums := []int{}
	n := len(nums1)
	k := k1 + k2

	for i := range nums1 {
		insertBinary(
			&nums,
			abs(max(nums1[i], nums2[i]) - min(nums1[i], nums2[i])),
		)
	}

	for k > 0 && nums[n - 1] > 0 {
		last := nums[n - 1]
		nums = nums[:n - 1]

		last--
		k--

		insertBinary(&nums, last)
	}

	var res int64
	for _, v := range nums {
		res += int64(v*v)
	}

	return res
}
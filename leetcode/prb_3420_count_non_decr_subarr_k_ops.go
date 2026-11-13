package leetcode


func countNonDecreasingSubarrays(nums []int, k int) int64 {
    // Steps = n (n - 1) / 2 => O(n^2) -- TLE 727/731
	// Space = O(1)

	var cnt int64

	for i := range nums {

		curK := k
		prev := 0
		for _, cur := range nums[i:] {
			if prev == 0 || cur >= prev {
				cnt++
				prev = cur
			} else if cur < prev && prev - cur <= curK {
				cnt++
				curK -= prev - cur
			} else {
				break
			}
		}
	}

	return cnt
}